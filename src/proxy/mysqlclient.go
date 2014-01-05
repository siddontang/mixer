package proxy

import (
	"encoding/binary"
	"errors"
	"github/siddontang/golib/log"
	"io"
	"net"
)

var (
	ErrMalformPacket   = errors.New("write Malform error")
	ErrPayloadLength   = errors.New("invalid payload length")
	ErrPacketSequence  = errors.New("invalid packet sequence")
	ErrProtocolVersion = errors.New("invalid protocol version")
)

//proxy <-> mysql server
//refer go-sql-driver
type MySQLClient struct {
	proxy      *Proxy
	address    string
	conn       *net.Conn
	sequence   uint8
	cipher     []byte
	capability uint32
	charset    uint8
	status     uint16
}

func NewMySQLClient(p *Proxy) *Client {
	c := new(Client)

	c.proxy = p

	return c
}

func (c *MySQLClient) Connect(address string) error {
	conn, err := net.Dial("tcp", address)
	if err {
		log.Error("connect %s error %s", address, err.Error())
		return err
	}

	if err := c.readInitPacket(); err != nil {
		log.Error("read initial handshake packet error %s", err.Error())
		return err
	}

	c.address = address
	c.conn = conn
	c.sequence = 0

	return nil
}

func (c *MySQLClient) readPacket() ([]byte, error) {
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
		buf, err = c.readPacket()
		if err != nil {
			log.Error("read packet error %s", err.Error())
			return nil, err
		} else {
			return append(data, buf)
		}
	}
}

//data already have header
func (c *MySQLClient) writePacket(data []byte) error {
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

func (c *MySQLClient) handleErrorPacket(data []byte) error {
	if data[0] != ERR_Packet {
		return ErrMalformPacket
	}

	errorCode := binary.LittleEndian.Uint16(data[1:3])

	pos := 3
	//sql state marker and state
	//maker is #, state is 5 length
	if data[pos] == '#' {
		pos = 9
	}

	return &MySQLError{
		Code:    errorCode,
		Message: data[pos:],
	}
}

func (c *MySQLClient) readInitPacket() error {
	data, err := c.readPacket()
	if err != nil {
		return err
	}

	if data[0] == ERR_Packet {
		return c.handleErrorPacket(data)
	}

	if data[0] < MinProtocolVersion {
		log.Error("invalid protocol version %d < %d", data[0], MinProtocolVersion)
		return ErrProtocolVersion
	}

	//skip mysql version and connection id
	//mysql version end with 0x00
	//connection id length is 4
	pos := 1 + bytes.IndexByte(data[1:], 0x00) + 1 + 4

	c.cipher = append(c.cipher, data[pos:pos+8]...)

	//skip filter
	pos += 8 + 1

	//capability lower 2 bytes
	c.capability = uint32(binary.LittleEndian.Uint16(data[pos : pos+2]))

	pos += 2

	if len(data) > pos {
		c.charset = data[pos]
		pos += 1
		c.status = binary.LittleEndian.Uint16(data[pos : pos+2])
	}
}
