package sqlparser

type Buffer interface {
	Write([]byte) (n, error)
	WriteString(string) (n, error)
}

type Formater interface {
	Format(b Buffer)
}

// Node represents basic type literal
type Node struct {
	Type  int
	Value string
}

func (n *Node) Format(b Buffer) {
	b.WriteString(n.Value)
}

func NewNode(tp int, value string) *Node {
	return &Node{tp, value}
}

type ParseError string

func (p ParseError) Error() string {
	return string(p)
}

func Parse(sql string) (Statement, error) {
	l := NewStringTokenizer(sql)
	if yyParse(l) != 0 {
		return nil, ParseError(l.LastError)
	}
	return l.ParseTree, nil
}
