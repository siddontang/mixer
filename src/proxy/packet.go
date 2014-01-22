package proxy

import (
	"encoding/binary"
	"errors"
)

var (
	ErrInvalidOKPacket  = errors.New("packet is not an ok packet")
	ErrInvalidErrPacket = errors.New("packet is not an error packet")
)

type OKPacket struct {
	AffectedRows uint64
	LastInsertId uint64
	Status       uint16
	Warnings     uint16
	Info         string
}

func DumpOK(pkg *OKPacket, capability uint32) []byte {
	data := make([]byte, 4, 32+len(pkg.Info))

	data = append(data, OK_HEADER)

	data = append(data, PutLengthEncodeInt(pkg.AffectedRows)...)
	data = append(data, PutLengthEncodeInt(pkg.LastInsertId)...)

	if capability|CLIENT_PROTOCOL_41 > 0 {
		data = append(data, byte(pkg.Status), byte(pkg.Status>>8))
		data = append(data, byte(pkg.Warnings), byte(pkg.Warnings>>8))
	} else if capability|CLIENT_TRANSACTIONS > 0 {
		data = append(data, byte(pkg.Status), byte(pkg.Status>>8))
	}

	data = append(data, pkg.Info...)

	return data
}

func DumpError(e error, capability uint32) []byte {
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

func LoadOK(data []byte, capability uint32) (*OKPacket, error) {
	if data[0] != OK_HEADER {
		return nil, ErrInvalidOKPacket
	}

	var n int
	var pos int = 1
	pkg := new(OKPacket)
	pkg.AffectedRows, _, n = LengthEncodeInt(data[pos:])
	pos += n
	pkg.LastInsertId, _, n = LengthEncodeInt(data[pos:])
	pos += n

	if capability&CLIENT_PROTOCOL_41 > 0 {
		pkg.Status = binary.LittleEndian.Uint16(data[pos:])
		pos += 2
		pkg.Warnings = binary.LittleEndian.Uint16(data[pos:])
		pos += 2
	} else if capability&CLIENT_TRANSACTIONS > 0 {
		pkg.Status = binary.LittleEndian.Uint16(data[pos:])
		pos += 2
	}

	pkg.Info = string(data[pos:])
	return pkg, nil
}

func LoadError(data []byte, capability uint32) (*MySQLError, error) {
	if data[0] != ERR_HEADER {
		return nil, ErrInvalidErrPacket
	}

	e := new(MySQLError)

	var pos int = 1

	e.Code = binary.LittleEndian.Uint16(data[pos:])
	pos += 2

	if capability&CLIENT_PROTOCOL_41 > 0 {
		//skip '#'
		pos++
		e.State = string(data[pos : pos+5])
		pos += 5
	}

	e.Message = string(data[pos:])

	return e, nil
}
