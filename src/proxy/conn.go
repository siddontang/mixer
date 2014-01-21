package proxy

import (
	"errors"
	"github.com/siddontang/golib/log"
	"io"
	"net"
)

var (
	ErrMalformPacket  = errors.New("write Malform error")
	ErrPayloadLength  = errors.New("invalid payload length")
	ErrPacketSequence = errors.New("invalid packet sequence")
)

type Conn struct {
	conn net.Conn

	sequence uint8
}

func (c *Conn) Close() error {
	return c.conn.Close()
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

func (c *Conn) BuildOK(affectedRows uint64, lastInsertId uint64,
	capability uint32, status uint16,
	warning uint16, info string) []byte {
	buf := make([]byte, 4, 32+len(info))

	buf = append(buf, OK_Packet)

	buf = append(buf, PutLengthEncodeInt(affectedRows)...)
	buf = append(buf, PutLengthEncodeInt(lastInsertId)...)

	if capability|CLIENT_PROTOCOL_41 > 0 {
		buf = append(buf, byte(status), byte(status>>8))
		buf = append(buf, byte(warning), byte(warning>>8))
	} else if capability|CLIENT_TRANSACTIONS > 0 {
		buf = append(buf, byte(status), byte(status>>8))
	}

	buf = append(buf, info...)

	return buf
}

func (c *Conn) BuildError(e error, capability uint32) []byte {
	var m *MySQLError
	var ok bool
	if m, ok = e.(*MySQLError); !ok {
		m = NewMySQLError(ER_UNKNOWN_ERROR, e.Error())
	}

	buf := make([]byte, 4, 16+len(m.Message))

	buf = append(buf, ERR_Packet)
	buf = append(buf, byte(m.Code), byte(m.Code>>8))

	buf = append(buf, '#')
	buf = append(buf, m.State...)

	buf = append(buf, m.Message...)

	return buf
}
