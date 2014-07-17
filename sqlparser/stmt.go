package sqlparser

type Stmt interface {
	Formater
	stmt()
}

type SelectStmt struct {
}

func (s *SelectStmt) stmt() {}

func (s *SelectStmt) Format(b Buffer) {

}

type UpdateStmt struct {
}

func (s *UpdateStmt) stmt() {}

func (s *UpdateStmt) Format(b Buffer) {

}

// DeleteStmt supports syntax:
//
//  delete from table [where where_condition] [order by ...] [limit row count]
type DeleteStmt struct {
	// 0x01 Low_PRIORITY
	// 0x02 QUICK
	// 0x04 IGNORE
	Opts  int
	Table *Node
	Where Expr
	Order Expr
	Limit Expr
}

func (s *DeleteStmt) stmt() {}

func (s *DeleteStmt) Format(b Buffer) {
	b.WriteString("DELETE ")
	if s.Opts&0x01 > 0 {
		b.WriteString("LOW_PRIORITY ")
	}
	if s.Opts&0x02 > 0 {
		b.WriteString("QUICK ")
	}

	if s.Opts&0x04 > 0 {
		b.WriteString("IGNORE ")
	}

	b.WriteString("FROM ")
	s.Table.Format(b)
}

type InsertStmt struct {
}

func (s *InsertStmt) stmt() {}

func (s *InsertStmt) Format(b Buffer) {

}

type ReplaceStmt struct {
}

func (s *ReplaceStmt) stmt() {}

func (s *ReplaceStmt) Format(b Buffer) {

}

type SetStmt struct {
}

func (s *SetStmt) stmt() {

}

func (s *SetStmt) Format(b Buffer) {

}
