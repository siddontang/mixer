package parser

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
refer:
    http://cuddle.googlecode.com/hg/talk/lex.html
    http://golang.org/src/pkg/text/template/parse/lex.go
*/

type stateFn func(*Lexer) stateFn

const eof = -1

type Lexer struct {
	input    string
	state    stateFn
	pos      int
	start    int
	width    int
	lastPos  int
	tokens   chan Token
	err      error
	stateMap map[rune]stateFn
}

func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	s := fmt.Sprintf(format, args...)
	l.err = fmt.Errorf("%s, near %d", s, l.start)
	return nil
}

func (l *Lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	return r
}

func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *Lexer) backup() {
	l.pos -= l.width
}

func (l *Lexer) emit(t TokenType) stateFn {
	l.tokens <- Token{t, l.input[l.start:l.pos]}
	l.start = l.pos
	return lexStart
}

func (l *Lexer) ignore() {
	l.start = l.pos
}

func (l *Lexer) reset() {
	l.pos = l.start
}

func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

func (l *Lexer) acceptRun(valid string) bool {
	var accepted bool = false
	for strings.IndexRune(valid, l.next()) >= 0 {
		accepted = true
	}
	l.backup()
	return accepted
}

func (l *Lexer) getToken() string {
	return l.input[l.start:l.pos]
}

func (l *Lexer) NextToken() (Token, error) {
	for {
		select {
		case t := <-l.tokens:
			return t, nil
		default:
			if l.state == nil {
				return Token{TK_EOF, ""}, l.err
			}
			l.state = l.state(l)
			if l.err != nil {
				return Token{TK_UNKNOWN, ""}, l.err
			}
		}
	}
}

func NewLexer(input string) *Lexer {
	l := new(Lexer)

	l.input = input
	l.tokens = make(chan Token, 2)
	l.state = lexStart
	l.stateMap = stateMap

	return l
}

var stateMap = map[rune]stateFn{
	',':  lexComma,
	';':  lexSemicolon,
	':':  lexColon,
	'.':  lexDot,
	'#':  lexLineComment,
	'"':  lexDoubleQuote,
	'\'': lexSignleQuote,
	'`':  lexRawQuote,
	'@':  lexVariable,
	'<':  lexLess,
	'>':  lexGreat,
	'=':  lexEqual,
	'!':  lexNot,
	'*':  lexStar,
	'+':  lexPlus,
	'/':  lexDiv,
	'-':  lexMinus,
	'&':  lexAnd,
	'|':  lexOr,
	'^':  lexXor,
	'%':  lexMod,
	'~':  lexInvert,
	'(':  lexLParent,
	')':  lexRParent,
	'[':  lexLBracket,
	']':  lexRBracket,
	'{':  lexLBrace,
	'}':  lexRBrace,
	'\\': lexBackslash,
	'0':  lexNumber0,
	'1':  lexNumber,
	'2':  lexNumber,
	'3':  lexNumber,
	'4':  lexNumber,
	'5':  lexNumber,
	'6':  lexNumber,
	'7':  lexNumber,
	'8':  lexNumber,
	'9':  lexNumber,
}

func lexStart(l *Lexer) stateFn {
	c := l.next()

	if c == eof {
		l.emit(TK_EOF)
		return nil
	} else if unicode.IsSpace(c) {
		return lexSpace
	} else if fn, ok := l.stateMap[c]; ok {
		return fn
	} else if isIdentifier(c) {
		return lexIdentifier
	}

	return l.errorf("unrecognized character in action: %v", c)
}

func lexSpace(l *Lexer) stateFn {
	for {
		c := l.next()
		if !unicode.IsSpace(c) {
			l.backup()
			break
		}
	}

	l.ignore()

	return lexStart
}

func scanNumber(l *Lexer) stateFn {
	digits := "0123456789"
	tokenType := TK_INTEGER

	l.acceptRun(digits)
	if l.accept(".") {
		tokenType = TK_FLOAT
		l.acceptRun(digits)
	}
	if l.accept("eE") {
		tokenType = TK_FLOAT
		l.accept("+-")
		l.acceptRun(digits)
	}

	return l.emit(tokenType)
}

func scanHex(l *Lexer) stateFn {
	l.acceptRun("0123456789abcdefABCDEF")

	return l.emit(TK_INTEGER)
}

func scanBin(l *Lexer) stateFn {
	l.acceptRun("01")

	return l.emit(TK_INTEGER)
}

func lexNumber0(l *Lexer) stateFn {
	c := l.next()
	switch c {
	case 'x', 'X':
		return scanHex(l)
	case 'b', 'B':
		return scanBin(l)
	default:
		l.backup()
		return scanNumber(l)
	}
}

func lexNumber(l *Lexer) stateFn {
	l.backup()

	return scanNumber(l)
}

func lexLineComment(l *Lexer) stateFn {
	i := strings.IndexByte(l.input[l.pos:], '\n')
	if i < 0 {
		//we assume no newline but eof
		l.pos = len(l.input)
		return l.emit(TK_COMMENT)
	} else {
		l.pos += i

		l.emit(TK_COMMENT)
		//skip '\n'
		l.pos += 1
		l.ignore()
	}

	return lexStart
}

func lexMultiLineComment(l *Lexer) stateFn {
	tp := TK_COMMENT
	if l.next() == '!' {
		tp = TK_COMMENT_MYSQL
	} else {
		l.backup()
	}

	i := strings.Index(l.input[l.pos:], "*/")
	if i < 0 {
		return l.errorf("unclosed comment")
	}

	l.pos += (i + 2)
	l.emit(tp)

	return lexStart
}

func scanQuote(l *Lexer, quote rune) bool {
Loop:
	for {
		switch l.next() {
		case eof:
			return false
		case '\\':
			if r := l.next(); r == eof {
				return false
			}
		case quote:
			if r := l.peek(); r == quote {
				l.next()
			} else {
				break Loop
			}
		}
	}

	return true
}

func lexQuote(l *Lexer, quote rune, tp TokenType) stateFn {
	if scanQuote(l, quote) {
		l.emit(tp)
		return lexStart
	} else {
		return l.errorf("unterminated quoted string")
	}
}

func lexDoubleQuote(l *Lexer) stateFn {
	return lexQuote(l, '"', TK_STRING)
}

func lexSignleQuote(l *Lexer) stateFn {
	return lexQuote(l, '\'', TK_STRING)
}

func lexRawQuote(l *Lexer) stateFn {
	return lexQuote(l, '`', TK_LITERAL)
}

func lexColon(l *Lexer) stateFn {
	c := l.next()
	if c == '=' {
		l.emit(TK_ASSIGN)
		return lexStart
	}

	return l.errorf("invalid character after ':' %v", c)
}

func lexSysVariable(l *Lexer) stateFn {
	switch r := l.next(); {
	case r == '`':
		return lexQuote(l, '`', TK_SYS_VARIABLE)
	case isVariable(r):
		for r = l.next(); isVariable(r); r = l.next() {
		}

		if r != eof {
			l.backup()
		}

		return l.emit(TK_SYS_VARIABLE)
	default:
		return l.errorf("invalid character %v", r)
	}
}

func lexVariable(l *Lexer) stateFn {
	//alphanumeric, ., _, $
	switch r := l.next(); {
	case r == '@':
		//system variable
		return lexSysVariable
	case r == '"':
		return lexQuote(l, '"', TK_VARIABLE)
	case r == '`':
		return lexQuote(l, '`', TK_VARIABLE)
	case r == '\'':
		return lexQuote(l, '\'', TK_VARIABLE)
	case isVariable(r):
		for r = l.next(); isVariable(r); r = l.next() {
		}

		if r != eof {
			l.backup()
		}

		return l.emit(TK_VARIABLE)
	default:
		return l.errorf("invalid character %#U", r)
	}
}

func lexIdentifier(l *Lexer) stateFn {
	//alphanumeric, _, $
	var r rune
	for r = l.next(); isIdentifier(r); r = l.next() {
	}

	if r != eof {
		l.backup()
	}

	//if ( followed by literal, we think literal is function
	if r == '(' {
		return l.emit(TK_FUNCTION)
	} else if r == '.' {
		//identifier may be seperated by .
		l.emit(getIdentiferType(l.getToken()))
		l.next()
		l.emit(TK_DOT)
		if isIdentifier(l.next()) {
			return lexIdentifier
		} else {
			l.backup()
			return lexStart
		}
	} else {
		return l.emit(getIdentiferType(l.getToken()))
	}
}

func lexDot(l *Lexer) stateFn {
	if c := l.peek(); unicode.IsNumber(c) {
		l.reset()
		return scanNumber(l)
	} else {
		l.emit(TK_DOT)
		return lexStart
	}
}

func lexComma(l *Lexer) stateFn {
	return l.emit(TK_COMMA)
}

func lexSemicolon(l *Lexer) stateFn {
	return l.emit(TK_SEMICOLON)
}

func lexLess(l *Lexer) stateFn {
	//<, <<, <=, <=>, <>
	switch c := l.next(); {
	case c == '<':
		return l.emit(TK_LTLT)
	case c == '>':
		return l.emit(TK_NE)
	case c == '=':
		if l.next() == '>' {
			//<=>
			return l.emit(TK_LTEQGT)
		} else {
			l.backup()
			return l.emit(TK_LE)
		}
	default:
		l.backup()
		return l.emit(TK_LT)
	}
}

func lexGreat(l *Lexer) stateFn {
	//>, >>, >=
	if r := l.next(); r == '>' || r == '=' {
		if r == '>' {
			return l.emit(TK_GTGT)
		} else {
			return l.emit(TK_GE)
		}
	} else {
		l.backup()
		return l.emit(TK_GT)
	}
}

func lexEqual(l *Lexer) stateFn {
	return l.emit(TK_EQ)
}

func lexNot(l *Lexer) stateFn {
	//!, !=
	if r := l.next(); r == '=' {
		return l.emit(TK_NE)
	} else {
		l.backup()
		return l.emit(TK_LOGICAL_NOT)
	}
}

func lexStar(l *Lexer) stateFn {
	return l.emit(TK_STAR)
}

func lexPlus(l *Lexer) stateFn {
	return l.emit(TK_PLUS)
}

func lexMinus(l *Lexer) stateFn {
	if l.next() == '-' && unicode.IsSpace(l.next()) {
		//comment "-- "
		return lexLineComment
	} else {
		l.reset()
		l.next()
		return l.emit(TK_MINUS)
	}
}

func lexDiv(l *Lexer) stateFn {
	if l.next() == '*' {
		return lexMultiLineComment
	} else {
		return l.emit(TK_DIV)
	}
}

func lexAnd(l *Lexer) stateFn {
	//&, &&
	if r := l.next(); r == '&' {
		return l.emit(TK_LOGICAL_AND)
	} else {
		l.backup()
		return l.emit(TK_BITWISE_AND)
	}

}

func lexOr(l *Lexer) stateFn {
	//|, ||
	if r := l.next(); r == '|' {
		return l.emit(TK_LOGICAL_OR)
	} else {
		l.backup()
		return l.emit(TK_BITWISE_OR)
	}
}

func lexXor(l *Lexer) stateFn {
	return l.emit(TK_BITWISE_XOR)
}

func lexMod(l *Lexer) stateFn {
	return l.emit(TK_MOD)
}

func lexInvert(l *Lexer) stateFn {
	return l.emit(TK_BITWISE_INVERT)
}

func lexLParent(l *Lexer) stateFn {
	return l.emit(TK_LPAREN)
}

func lexRParent(l *Lexer) stateFn {
	return l.emit(TK_RPAREN)
}

func lexLBracket(l *Lexer) stateFn {
	return l.emit(TK_LBRACKET)
}

func lexRBracket(l *Lexer) stateFn {
	return l.emit(TK_RBRACKET)
}

func lexLBrace(l *Lexer) stateFn {
	return l.emit(TK_LBRACE)
}

func lexRBrace(l *Lexer) stateFn {
	return l.emit(TK_RBRACE)
}

func lexBackslash(l *Lexer) stateFn {
	if l.next() == 'N' {
		//shortcut for NULL
		return l.emit(TK_SQL_NULL)
	} else {
		l.backup()
		return l.emit(TK_BACKSLASH)
	}
}

func isIdentifier(r rune) bool {
	return unicode.IsNumber(r) || unicode.IsLetter(r) || r == '$' || r == '_'
}

func isVariable(r rune) bool {
	return unicode.IsNumber(r) || unicode.IsLetter(r) || r == '$' || r == '_' || r == '.'
}

//get identifier type
func getIdentiferType(str string) TokenType {
	if tp, ok := KeyWords[strings.ToUpper(str)]; ok {
		return tp
	} else {
		return TK_LITERAL
	}
}

func Tokenizer(input string) ([]Token, error) {
	l := NewLexer(input)

	ts := make([]Token, 0)

	for {
		t, err := l.NextToken()
		if err != nil {
			return nil, err
		}

		if t.Type == TK_EOF {
			break
		}

		ts = append(ts, t)
	}

	return ts, nil
}

var (
	errLogger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
)

func errLog(format string, args ...interface{}) {
	f := fmt.Sprintf("[Error] [mixer.parser] %s", format)
	s := fmt.Sprintf(f, args...)
	errLogger.Output(2, s)
}
