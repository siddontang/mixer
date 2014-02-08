package mysql

import (
	"database/sql/driver"
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

type PacketIO struct {
	Conn     net.Conn
	Sequence uint8
}

func (c *PacketIO) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *PacketIO) LocalAddr() net.Addr {
	return c.Conn.LocalAddr()
}

func (c *PacketIO) ReadPacket() ([]byte, error) {
	header := make([]byte, 4)

	if _, err := io.ReadFull(c.Conn, header); err != nil {
		log.Error("read header error %s", err.Error())
		return nil, driver.ErrBadConn
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
	if _, err := io.ReadFull(c.Conn, data); err != nil {
		log.Error("read payload data error %s", err.Error())
		return nil, driver.ErrBadConn
	} else {
		if length < MaxPayloadLen {
			return data, nil
		}

		var buf []byte
		buf, err = c.ReadPacket()
		if err != nil {
			log.Error("read packet error %s", err.Error())
			return nil, driver.ErrBadConn
		} else {
			return append(data, buf...), nil
		}
	}
}

//data already have header
func (c *PacketIO) WritePacket(data []byte) error {
	length := len(data) - 4

	for length >= MaxPayloadLen {

		data[0] = 0xff
		data[1] = 0xff
		data[2] = 0xff

		data[3] = c.Sequence

		if n, err := c.Conn.Write(data[:4+MaxPayloadLen]); err != nil {
			log.Error("write error %s", err.Error())
			return driver.ErrBadConn
		} else if n != (4 + MaxPayloadLen) {
			log.Error("write error, write data number %d != %d", n, (4 + MaxPayloadLen))
			return driver.ErrBadConn
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

	if n, err := c.Conn.Write(data); err != nil {
		log.Error("write error %s", err.Error())
		return driver.ErrBadConn
	} else if n != len(data) {
		log.Error("write error, write data number %d != %d", n, (4 + MaxPayloadLen))
		return driver.ErrBadConn
	} else {
		c.Sequence++
		return nil
	}
}

func (c *PacketIO) WriteOK(pkg *OKPacket) error {
	data := DumpOK(pkg)

	return c.WritePacket(data)
}

func (c *PacketIO) WriteError(e error) error {
	data := DumpError(e)

	return c.WritePacket(data)
}

func (c *PacketIO) WriteEOF(pkg *EOFPacket) error {
	data := DumpEOF(pkg)

	return c.WritePacket(data)
}

func (c *PacketIO) ReadOK() (*OKPacket, error) {
	data, err := c.ReadPacket()
	if err != nil {
		return nil, err
	}

	if data[0] == OK_HEADER {
		return LoadOK(data), nil
	} else if data[0] == ERR_HEADER {
		return nil, LoadError(data)
	} else {
		return nil, ErrInvalidOKPacket
	}
}
