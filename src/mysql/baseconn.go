package mysql

import (
	"encoding/binary"
	"errors"
	"github.com/siddontang/golib/log"
	"io"
	"net"
)

var (
	ErrMalformPacket    = errors.New("Malform packet error")
	ErrPayloadLength    = errors.New("Invalid payload length")
	ErrPacketSequence   = errors.New("Invalid packet sequence")
	ErrInvalidOKPacket  = errors.New("Packet is not an ok packet")
	ErrInvalidErrPacket = errors.New("Packet is not an error packet")
)

type BaseConn struct {
	Connection net.Conn

	Sequence uint8

	Capability uint32
}

func (c *BaseConn) Close() error {
	err := c.Connection.Close()
	c.Connection = nil
	return err
}

func (c *BaseConn) RemoteAddr() net.Addr {
	return c.Connection.RemoteAddr()
}

func (c *BaseConn) LocalAddr() net.Addr {
	return c.Connection.LocalAddr()
}

func (c *BaseConn) ReadPacket() ([]byte, error) {
	header := make([]byte, 4)

	if _, err := io.ReadFull(c.Connection, header); err != nil {
		return nil, err
	}

	length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)
	if length < 1 {
		log.Error("invalid payload length")
		return nil, ErrPayloadLength
	}

	sequence := uint8(header[3])

	if sequence != c.Sequence {
		log.Error("invalid sequence %d != %d", sequence, c.Sequence)
		return nil, ErrPacketSequence
	}

	c.Sequence++

	data := make([]byte, length)
	if _, err := io.ReadFull(c.Connection, data); err != nil {
		log.Error("read payload data error %s", err.Error())
		return nil, err
	} else {
		if length < MaxPayloadLen {
			return data, nil
		}

		var buf []byte
		buf, err = c.ReadPacket()
		if err != nil {
			log.Error("read packet error %s", err.Error())
			return nil, err
		} else {
			return append(data, buf...), nil
		}
	}
}

//data already have header
func (c *BaseConn) WritePacket(data []byte) error {
	length := len(data) - 4

	for length >= MaxPayloadLen {

		data[0] = 0xff
		data[1] = 0xff
		data[2] = 0xff

		data[3] = c.Sequence

		if n, err := c.Connection.Write(data[:4+MaxPayloadLen]); err != nil {
			log.Error("write error %s", err.Error())
			return err
		} else if n != (4 + MaxPayloadLen) {
			log.Error("write error, write data number %d != %d", n, (4 + MaxPayloadLen))
			return ErrMalformPacket
		} else {
			c.Sequence++
			length -= MaxPayloadLen
			data = data[MaxPayloadLen:]
		}
	}

	data[0] = byte(length)
	data[1] = byte(length >> 8)
	data[2] = byte(length >> 16)
	data[3] = c.Sequence

	if n, err := c.Connection.Write(data); err != nil {
		log.Error("write error %s", err.Error())
		return err
	} else if n != len(data) {
		log.Error("write error, write data number %d != %d", n, (4 + MaxPayloadLen))
		return ErrMalformPacket
	} else {
		c.Sequence++
		return nil
	}
}

func (c *BaseConn) DumpOK(pkg *OKPacket) []byte {
	data := make([]byte, 4, 32+len(pkg.Info))

	data = append(data, OK_HEADER)

	data = append(data, PutLengthEncodedInt(pkg.AffectedRows)...)
	data = append(data, PutLengthEncodedInt(pkg.LastInsertId)...)

	if c.Capability|CLIENT_PROTOCOL_41 > 0 {
		data = append(data, byte(pkg.Status), byte(pkg.Status>>8))
		data = append(data, byte(pkg.Warnings), byte(pkg.Warnings>>8))
	} else if c.Capability|CLIENT_TRANSACTIONS > 0 {
		data = append(data, byte(pkg.Status), byte(pkg.Status>>8))
	}

	data = append(data, pkg.Info...)

	return data
}

func (c *BaseConn) DumpError(e error) []byte {
	var m *MySQLError
	var ok bool
	if m, ok = e.(*MySQLError); !ok {
		m = NewError(ER_UNKNOWN_ERROR, e.Error())
	}

	data := make([]byte, 4, 16+len(m.Message))

	data = append(data, ERR_HEADER)
	data = append(data, byte(m.Code), byte(m.Code>>8))

	data = append(data, '#')
	data = append(data, m.State...)

	data = append(data, m.Message...)

	return data
}

func (c *BaseConn) DumpEOF(pkg *EOFPacket) []byte {
	data := make([]byte, 4, 9)

	data = append(data, EOF_HEADER)

	if c.Capability&CLIENT_PROTOCOL_41 > 0 {
		data = append(data, byte(pkg.Warnings), byte(pkg.Warnings>>8))
		data = append(data, byte(pkg.Status), byte(pkg.Status>>8))
	}

	return data
}

func (c *BaseConn) LoadOK(data []byte) *OKPacket {
	if data[0] != OK_HEADER {
		return nil
	}

	var n int
	var pos int = 1

	pkg := new(OKPacket)
	pkg.AffectedRows, _, n = LengthEncodedInt(data[pos:])
	pos += n
	pkg.LastInsertId, _, n = LengthEncodedInt(data[pos:])
	pos += n

	if c.Capability&CLIENT_PROTOCOL_41 > 0 {
		pkg.Status = binary.LittleEndian.Uint16(data[pos:])
		pos += 2
		pkg.Warnings = binary.LittleEndian.Uint16(data[pos:])
		pos += 2
	} else if c.Capability&CLIENT_TRANSACTIONS > 0 {
		pkg.Status = binary.LittleEndian.Uint16(data[pos:])
		pos += 2
	}

	pkg.Info = string(data[pos:])
	return pkg
}

func (c *BaseConn) LoadError(data []byte) *MySQLError {
	if data[0] != ERR_HEADER {
		return nil
	}

	e := new(MySQLError)

	var pos int = 1

	e.Code = binary.LittleEndian.Uint16(data[pos:])
	pos += 2

	if c.Capability&CLIENT_PROTOCOL_41 > 0 {
		//skip '#'
		pos++
		e.State = string(data[pos : pos+5])
		pos += 5
	}

	e.Message = string(data[pos:])

	return e
}

func (c *BaseConn) LoadEOF(data []byte) *EOFPacket {
	if data[0] != EOF_HEADER || len(data) > 5 {
		//length encoded int may begin with 0xfe too
		return nil
	}

	pkg := new(EOFPacket)
	if c.Capability&CLIENT_PROTOCOL_41 > 0 {
		pkg.Warnings = binary.LittleEndian.Uint16(data[1:])
		pkg.Status = binary.LittleEndian.Uint16(data[3:])
	}

	return pkg
}

func (c *BaseConn) WriteOK(pkg *OKPacket) error {
	data := c.DumpOK(pkg)

	return c.WritePacket(data)
}

func (c *BaseConn) WriteError(e error) error {
	data := c.DumpError(e)

	return c.WritePacket(data)
}

func (c *BaseConn) WriteEOF(pkg *EOFPacket) error {
	data := c.DumpEOF(pkg)

	return c.WritePacket(data)
}

func (c *BaseConn) ReadOK() (*OKPacket, error) {
	data, err := c.ReadPacket()
	if err != nil {
		return nil, err
	}

	if data[0] == OK_HEADER {
		return c.LoadOK(data), nil
	} else if data[0] == ERR_HEADER {
		return nil, c.LoadError(data)
	} else {
		return nil, ErrInvalidOKPacket
	}
}
