package sqlparser

type Expr interface {
	Formater
	expr()
}
