package sqlparser

type Node struct {
}

func (n *Node) Format(b *Buffer) {

}

type Statement interface {
	statement()
	Format(*Buffer)
}

type Select struct {
}

func (s *Select) statement() {}

func (s *Select) Format(b *Buffer) {

}
