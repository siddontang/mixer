package mysql

import (
	"encoding/binary"
)

type ColumnField struct {
	Name string
	Type byte
	Flag uint16
}

type ResultSet struct {
	Columns []ColumnField

	Rows [][]interface{}
}

func NewTextResultSet(columns [][]byte, rows [][]byte) (*ResultSet, error) {
	r := new(ResultSet)

	if err := r.handleTextColumns(columns); err != nil {
		return nil, err
	}

	if err := r.handleTextRows(rows); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *ResultSet) handleTextColumns(columns [][]byte) (err error) {
	r.Columns = make([]ColumnField, len(columns))

	var pos int
	var n int

	for i := 0; i < len(columns); i++ {
		data := columns[i]

		// Catalog
		pos, err = skipLengthEnodedString(data)
		if err != nil {
			return
		}

		// Database [len coded string]
		n, err = skipLengthEnodedString(data[pos:])
		if err != nil {
			return
		}
		pos += n

		// Table [len coded string]
		n, err = skipLengthEnodedString(data[pos:])
		if err != nil {
			return
		}
		pos += n

		// Original table [len coded string]
		n, err = skipLengthEnodedString(data[pos:])
		if err != nil {
			return
		}
		pos += n

		// Name [len coded string]
		name, _, n, err = readLengthEnodedString(data[pos:])
		if err != nil {
			return
		}
		r.Columns[i].Name = string(name)
		pos += n

		// Original name [len coded string]
		n, err = skipLengthEnodedString(data[pos:])
		if err != nil {
			return
		}

		// Filler [1 byte]
		// Charset [16 bit uint]
		// Length [32 bit uint]
		pos += n + 1 + 2 + 4

		// Field type [byte]
		r.Columns[i].Type = data[pos]
		pos++

		// Flags [16 bit uint]
		r.Columns[i].Flag = binary.LittleEndian.Uint16(data[pos : pos+2])

		// skip left
	}

	return nil
}

func (r *ResultSet) handleTextRows(rows [][]byte) (err error) {
	r.Rows = make([][]interface{}, len(rows))

	var isNull bool
	var n int
	var item []byte

	for i := 0; i < len(rows); i++ {
		r.Rows[i] = make([]interface{}, len(r.Columns))

		pos := 0

		data := rows[i]

		for j := range r.Rows[i] {
			item, isNull, n, err = readLengthEnodedString(data[pos:])
			pos += n

			if err != nil {
				return
			}

			if isNull {
				r.Rows[i][j] = nil
			} else {
				r.Rows[i][j] = item
			}
		}

	}

	return nil
}
