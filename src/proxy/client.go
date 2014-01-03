package proxy

import (
	"errors"
	"github/siddontang/golib/log"
	"io"
	"net"
)

var (
	ErrPayloadLength  = errors.New("invalid payload length")
	ErrPacketSequence = errors.New("invalid packet sequence")
)

//proxy <-> mysql server
type Client struct {
	proxy    *Proxy
	address  string
	conn     *net.Conn
	sequence uint8
}

func NewClient(p *Proxy) *Client {
	c := new(Client)

	c.proxy = p

	return c
}

func (c *Client) Connect(address string) error {
	conn, err := net.Dial("tcp", address)
	if err {
		log.Error("connect %s error %s", address, err.Error())
		return err
	}

	if err := c.onHandshake(); err != nil {
		log.Error("handshake error %s", err.Error())
		return err
	}

	c.address = address
	c.conn = conn
	c.sequence = 0

	return nil
}

func (c *Client) onHandshake() error {

}

func (c *Client) readPackets() ([]byte, error) {
	head := make([]byte, 4)

	if _, err := io.ReadFull(c.conn, head); err != nil {
		return nil, err
	}

	length := int(uint32(head[0]) | uint32(head[1])<<8 | uint32(head[2])<<16)
	if length < 1 {
		log.Error("invalid payload length")
		return nil, ErrPayloadLength
	}

	sequence := uint8(head[3])

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
		buf, err = c.readPackets()
		if err != nil {
			log.Error("read packet error %s", err.Error())
			return nil, err
		} else {
			return append(data, buf)
		}
	}
}

func (c *Client) writePackets(data []byte) error {
	length := len(data)

	if length >= MaxPayloadLen {

	}
}
