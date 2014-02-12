package mysql

import (
	"encoding/binary"
)

type OKPacket struct {
	AffectedRows uint64
	LastInsertId uint64
	Status       uint16
	Warnings     uint16
	Info         string
}

type EOFPacket struct {
	Status   uint16
	Warnings uint16
}

func DumpOK(pkg *OKPacket) []byte {
	data := make([]byte, 4, 32+len(pkg.Info))

	data = append(data, OK_HEADER)

	data = append(data, PutLengthEncodedInt(pkg.AffectedRows)...)
	data = append(data, PutLengthEncodedInt(pkg.LastInsertId)...)

	data = append(data, byte(pkg.Status), byte(pkg.Status>>8))
	data = append(data, byte(pkg.Warnings), byte(pkg.Warnings>>8))

	// if capability|CLIENT_PROTOCOL_41 > 0 {
	//    data = append(data, byte(pkg.Status), byte(pkg.Status>>8))
	//    data = append(data, byte(pkg.Warnings), byte(pkg.Warnings>>8))
	// } else if capability|CLIENT_TRANSACTIONS > 0 {
	//    data = append(data, byte(pkg.Status), byte(pkg.Status>>8))
	// }

	data = append(data, pkg.Info...)

	return data
}

func DumpError(e error) []byte {
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

func DumpEOF(pkg *EOFPacket) []byte {
	data := make([]byte, 4, 9)

	data = append(data, EOF_HEADER)

	data = append(data, byte(pkg.Warnings), byte(pkg.Warnings>>8))
	data = append(data, byte(pkg.Status), byte(pkg.Status>>8))

	// if c.Capability&CLIENT_PROTOCOL_41 > 0 {
	// 	data = append(data, byte(pkg.Warnings), byte(pkg.Warnings>>8))
	// 	data = append(data, byte(pkg.Status), byte(pkg.Status>>8))
	// }

	return data
}

func LoadOK(data []byte) *OKPacket {
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

	pkg.Status = binary.LittleEndian.Uint16(data[pos:])
	pos += 2
	pkg.Warnings = binary.LittleEndian.Uint16(data[pos:])
	pos += 2

	// if c.Capability&CLIENT_PROTOCOL_41 > 0 {
	// 	pkg.Status = binary.LittleEndian.Uint16(data[pos:])
	// 	pos += 2
	// 	pkg.Warnings = binary.LittleEndian.Uint16(data[pos:])
	// 	pos += 2
	// } else if c.Capability&CLIENT_TRANSACTIONS > 0 {
	// 	pkg.Status = binary.LittleEndian.Uint16(data[pos:])
	// 	pos += 2
	// }

	pkg.Info = string(data[pos:])
	return pkg
}

func LoadError(data []byte) *MySQLError {
	if data[0] != ERR_HEADER {
		return nil
	}

	e := new(MySQLError)

	var pos int = 1

	e.Code = binary.LittleEndian.Uint16(data[pos:])
	pos += 2

	//skip '#'
	pos++
	e.State = string(data[pos : pos+5])
	pos += 5

	// if c.Capability&CLIENT_PROTOCOL_41 > 0 {
	// 	//skip '#'
	// 	pos++
	// 	e.State = string(data[pos : pos+5])
	// 	pos += 5
	// }

	e.Message = string(data[pos:])

	return e
}

func LoadEOF(data []byte) *EOFPacket {
	if data[0] != EOF_HEADER || len(data) > 5 {
		//length encoded int may begin with 0xfe too
		return nil
	}

	pkg := new(EOFPacket)
	pkg.Warnings = binary.LittleEndian.Uint16(data[1:])
	pkg.Status = binary.LittleEndian.Uint16(data[3:])

	// if c.Capability&CLIENT_PROTOCOL_41 > 0 {
	// 	pkg.Warnings = binary.LittleEndian.Uint16(data[1:])
	// 	pkg.Status = binary.LittleEndian.Uint16(data[3:])
	// }

	return pkg
}
