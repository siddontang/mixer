package proxy

import (
	"encoding/binary"
	"errors"
	"github.com/siddontang/golib/log"
	"io"
	"net"
)

var (
	ErrMalformPacket    = errors.New("Malform packet error")
	ErrPayloadLength    = errors.New("invalid payload length")
	ErrPacketSequence   = errors.New("invalid packet sequence")
	ErrInvalidOKPacket  = errors.New("packet is not an ok packet")
	ErrInvalidErrPacket = errors.New("packet is not an error packet")
)

type Conn struct {
	conn net.Conn

	sequence uint8

	capability uint32
}

func (c *Conn) Close() error {
	err := c.conn.Close()
	c.conn = nil
	return err
}

func (c *Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *Conn) ReadPacket() ([]byte, error) {
	header := make([]byte, 4)

	if _, err := io.ReadFull(c.conn, header); err != nil {
		return nil, err
	}

	length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)
	if length < 1 {
		log.Error("invalid payload length")
		return nil, ErrPayloadLength
	}

	sequence := uint8(header[3])

	if sequence != c.sequence {
		log.Error("invalid sequence %d != %d", sequence, c.sequence)
		return nil, ErrPacketSequence
	}

	c.sequence++

	data := make([]byte, length)
	if _, err := io.ReadFull(c.conn, data); err != nil {
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
func (c *Conn) WritePacket(data []byte) error {
	length := len(data) - 4

	for length >= MaxPayloadLen {

		data[0] = 0xff
		data[1] = 0xff
		data[2] = 0xff

		data[3] = c.sequence

		if n, err := c.conn.Write(data[:4+MaxPayloadLen]); err != nil {
			log.Error("write error %s", err.Error())
			return err
		} else if n != (4 + MaxPayloadLen) {
			log.Error("write error, write data number %d != %d", n, (4 + MaxPayloadLen))
			return ErrMalformPacket
		} else {
			c.sequence++
			length -= MaxPayloadLen
			data = data[MaxPayloadLen:]
		}
	}

	data[0] = byte(length)
	data[1] = byte(length >> 8)
	data[2] = byte(length >> 16)
	data[3] = c.sequence

	if n, err := c.conn.Write(data); err != nil {
		log.Error("write error %s", err.Error())
		return err
	} else if n != len(data) {
		log.Error("write error, write data number %d != %d", n, (4 + MaxPayloadLen))
		return ErrMalformPacket
	} else {
		c.sequence++
		return nil
	}
}

func (c *Conn) DumpOK(pkg *OKPacket) []byte {
	data := make([]byte, 4, 32+len(pkg.Info))

	data = append(data, OK_HEADER)

	data = append(data, PutLengthEncodedInt(pkg.AffectedRows)...)
	data = append(data, PutLengthEncodedInt(pkg.LastInsertId)...)

	if c.capability|CLIENT_PROTOCOL_41 > 0 {
		data = append(data, byte(pkg.Status), byte(pkg.Status>>8))
		data = append(data, byte(pkg.Warnings), byte(pkg.Warnings>>8))
	} else if c.capability|CLIENT_TRANSACTIONS > 0 {
		data = append(data, byte(pkg.Status), byte(pkg.Status>>8))
	}

	data = append(data, pkg.Info...)

	return data
}

func (c *Conn) DumpError(e error) []byte {
	var m *MySQLError
	var ok bool
	if m, ok = e.(*MySQLError); !ok {
		m = NewMySQLError(ER_UNKNOWN_ERROR, e.Error())
	}

	data := make([]byte, 4, 16+len(m.Message))

	data = append(data, ERR_HEADER)
	data = append(data, byte(m.Code), byte(m.Code>>8))

	data = append(data, '#')
	data = append(data, m.State...)

	data = append(data, m.Message...)

	return data
}

func (c *Conn) DumpEOF(pkg *EOFPacket) []byte {
	data := make([]byte, 4, 8)

	if c.capability&CLIENT_PROTOCOL_41 > 0 {
		data = append(data, byte(pkg.Warnings), byte(pkg.Warnings>>8))
		data = append(data, byte(pkg.Status), byte(pkg.Status>>8))
	}

	return data
}

func (c *Conn) LoadOK(data []byte) *OKPacket {
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

	if c.capability&CLIENT_PROTOCOL_41 > 0 {
		pkg.Status = binary.LittleEndian.Uint16(data[pos:])
		pos += 2
		pkg.Warnings = binary.LittleEndian.Uint16(data[pos:])
		pos += 2
	} else if c.capability&CLIENT_TRANSACTIONS > 0 {
		pkg.Status = binary.LittleEndian.Uint16(data[pos:])
		pos += 2
	}

	pkg.Info = string(data[pos:])
	return pkg
}

func (c *Conn) LoadError(data []byte) *MySQLError {
	if data[0] != ERR_HEADER {
		return nil
	}

	e := new(MySQLError)

	var pos int = 1

	e.Code = binary.LittleEndian.Uint16(data[pos:])
	pos += 2

	if c.capability&CLIENT_PROTOCOL_41 > 0 {
		//skip '#'
		pos++
		e.State = string(data[pos : pos+5])
		pos += 5
	}

	e.Message = string(data[pos:])

	return e
}

func (c *Conn) LoadEOF(data []byte) *EOFPacket {
	if data[0] != EOF_HEADER || len(data) > 5 {
		//length encoded int may begin with 0xfe too
		return nil
	}

	pkg := new(EOFPacket)
	if c.capability&CLIENT_PROTOCOL_41 > 0 {
		pkg.Warnings = binary.LittleEndian.Uint16(data[1:])
		pkg.Status = binary.LittleEndian.Uint16(data[3:])
	}

	return pkg
}

func (c *Conn) WriteOK(pkg *OKPacket) error {
	data := c.DumpOK(pkg)

	return c.WritePacket(data)
}

func (c *Conn) WriteError(e error) error {
	data := c.DumpError(e)

	return c.WritePacket(data)
}

func (c *Conn) WriteEOF(pkg *EOFPacket) error {
	data := c.DumpEOF(pkg)

	return c.WritePacket(data)
}

func (c *Conn) ReadOK() (*OKPacket, error) {
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
