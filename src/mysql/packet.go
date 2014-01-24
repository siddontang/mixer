package mysql

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

type ColumnDefPacket struct {
	Schema   string
	Table    string
	OrgTable string
	Name     string
	OrgName  string
	Charset  uint16
	Length   uint32
	Type     uint8
	Flag     uint16
	Decimals uint8

	//below if command was COM_FIELD_LIST
	DefaultLen   uint64
	DefaultValue string
}

type TextResultPacket struct {
	ColumnDefs [][]byte
	Rows       [][]byte
}
