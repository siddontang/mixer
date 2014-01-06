package mysql

import ()

type RawResultSet struct {
	columns [][]byte
	rows    [][]byte
}

type ColumnField struct {
	Name string
	Type byte
	Flag uint16
}

type ResultSet struct {
	Columns []ColumnField

	Rows [][]interface{}
}
