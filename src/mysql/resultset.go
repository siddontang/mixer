package mysql

import (
	"encoding/binary"
	"errors"
	"math"
	"strconv"
)

var (
	ErrInvalidFieldType = errors.New("invalid field type")
	ErrInvalidIndex     = errors.New("invalid index")
)

type ColumnField struct {
	Name string
	Type byte
	Flag uint16
}

type ResultSet struct {
	Columns []ColumnField

	ColumnNames map[string]int

	Rows [][]interface{}
}

func (r *ResultSet) GetInt(row, column int) (int64, error) {
	if err := r.checkIndex(row, column); err != nil {
		return 0, err
	}

	switch v := r.Rows[row][column].(type) {
	case int64:
		return v, nil
	case uint64:
		return int64(v), nil
	default:
		return int64(0), nil
	}
}

func (r *ResultSet) GetUInt(row, column int) (uint64, error) {
	if err := r.checkIndex(row, column); err != nil {
		return 0, err
	}

	switch v := r.Rows[row][column].(type) {
	case int64:
		return uint64(v), nil
	case uint64:
		return v, nil
	default:
		return uint64(0), nil
	}
}

func (r *ResultSet) GetFloat(row, column int) (float64, error) {
	if err := r.checkIndex(row, column); err != nil {
		return 0, err
	}

	switch v := r.Rows[row][column].(type) {
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case int64:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	default:
		return float64(0), nil
	}
}

func (r *ResultSet) GetBool(row, column int) (bool, error) {
	if err := r.checkIndex(row, column); err != nil {
		return false, err
	}

	switch v := r.Rows[row][column].(type) {
	case bool:
		return v, nil
	case int64:
		return (v != int64(0)), nil
	case uint64:
		return (v != uint64(0)), nil
	default:
		return false, nil
	}
}

func (r *ResultSet) GetString(row, column int) (string, error) {
	if err := r.checkIndex(row, column); err != nil {
		return "", err
	}

	switch v := r.Rows[row][column].(type) {
	case []byte:
		return string(v), nil
	case string:
		return v, nil
	default:
		return "", nil
	}
}

func (r *ResultSet) GetBytes(row, column int) ([]byte, error) {
	if err := r.checkIndex(row, column); err != nil {
		return nil, err
	}

	switch v := r.Rows[row][column].(type) {
	case []byte:
		return v, nil
	case string:
		return []byte(v), nil
	default:
		return nil, nil
	}
}

func (r *ResultSet) GetIntByName(row int, colName string) (int64, error) {
	if col, ok := r.ColumnNames[colName]; ok {
		return r.GetInt(row, col)
	} else {
		return 0, ErrInvalidIndex
	}
}

func (r *ResultSet) GetUIntByName(row int, colName string) (uint64, error) {
	if col, ok := r.ColumnNames[colName]; ok {
		return r.GetUInt(row, col)
	} else {
		return 0, ErrInvalidIndex
	}
}

func (r *ResultSet) GetFloatByName(row int, colName string) (float64, error) {
	if col, ok := r.ColumnNames[colName]; ok {
		return r.GetFloat(row, col)
	} else {
		return float64(0), ErrInvalidIndex
	}
}

func (r *ResultSet) GetBoolByName(row int, colName string) (bool, error) {
	if col, ok := r.ColumnNames[colName]; ok {
		return r.GetBool(row, col)
	} else {
		return false, ErrInvalidIndex
	}
}

func (r *ResultSet) GetStringByName(row int, colName string) (string, error) {
	if col, ok := r.ColumnNames[colName]; ok {
		return r.GetString(row, col)
	} else {
		return "", ErrInvalidIndex
	}
}

func (r *ResultSet) GetBytesByName(row int, colName string) ([]byte, error) {
	if col, ok := r.ColumnNames[colName]; ok {
		return r.GetBytes(row, col)
	} else {
		return nil, ErrInvalidIndex
	}
}

func (r *ResultSet) checkIndex(row, column int) error {
	if row >= len(r.Rows) || column >= len(r.Columns) {
		return ErrInvalidIndex
	}

	return nil
}

func (r *ResultSet) handleTextColumns(columns [][]byte) (err error) {
	r.Columns = make([]ColumnField, len(columns))
	r.ColumnNames = make(map[string]int, len(columns))

	var pos int
	var n int
	var name []byte

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
		r.ColumnNames[r.Columns[i].Name] = i

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
				switch r.Columns[j].Type {
				case MYSQL_TYPE_NULL:
					{
						r.Rows[i][j] = nil
					}
				case MYSQL_TYPE_TINY, MYSQL_TYPE_SHORT, MYSQL_TYPE_LONG, MYSQL_TYPE_INT24, MYSQL_TYPE_YEAR:
					{
						r.Rows[i][j], _ = strconv.ParseInt(string(item), 10, 64)
					}
				case MYSQL_TYPE_LONGLONG:
					{
						v, _ := strconv.ParseUint(string(item), 10, 64)
						if v <= math.MaxInt64 {
							r.Rows[i][j] = int64(v)
						} else {
							r.Rows[i][j] = v
						}
					}
				case MYSQL_TYPE_FLOAT, MYSQL_TYPE_DOUBLE:
					{
						r.Rows[i][j], _ = strconv.ParseFloat(string(item), 64)
					}
				case MYSQL_TYPE_DECIMAL,
					MYSQL_TYPE_TIMESTAMP,
					MYSQL_TYPE_DATE,
					MYSQL_TYPE_TIME,
					MYSQL_TYPE_DATETIME,
					MYSQL_TYPE_NEWDATE,
					MYSQL_TYPE_VARCHAR,
					MYSQL_TYPE_BIT,
					MYSQL_TYPE_NEWDECIMAL,
					MYSQL_TYPE_ENUM,
					MYSQL_TYPE_SET,
					MYSQL_TYPE_TINY_BLOB,
					MYSQL_TYPE_MEDIUM_BLOB,
					MYSQL_TYPE_LONG_BLOB,
					MYSQL_TYPE_BLOB,
					MYSQL_TYPE_VAR_STRING,
					MYSQL_TYPE_STRING,
					MYSQL_TYPE_GEOMETRY:
					{
						r.Rows[i][j] = item
					}
				default:
					return ErrInvalidFieldType
				}
			}
		}

	}

	return nil
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
