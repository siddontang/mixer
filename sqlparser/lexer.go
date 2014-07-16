package sqlparser

type Lexer struct {
	ParseTree Statement
}

func (l *Lexer) Lex(lval *yySymType) int {

}

func (l *Lexer) Error(e string) {

}
