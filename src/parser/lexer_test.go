package parser

import (
	"fmt"
	"testing"
)

func testCheckTokens(tokens []Token, expected []Token) error {
	if len(tokens) != len(expected) {
		return fmt.Errorf("expect %d tokens, but %d", len(expected), len(tokens))
	}

	for i := range tokens {
		if tokens[i].Type != expected[i].Type {
			return fmt.Errorf("%d invalid type %d != %d", i, tokens[i].Type, expected[i].Type)
		}

		if tokens[i].Value != expected[i].Value {
			return fmt.Errorf("%d invalid value %s(%d) != %s(%d)", i, tokens[i].Value, len(tokens[i].Value), expected[i].Value, len(expected[i].Value))

		}
	}

	return nil
}

func TestLexer_String(t *testing.T) {
	s := "\"hello world\""

	tokens, err := Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	if err := testCheckTokens(tokens, []Token{
		Token{TK_STRING, "\"hello world\""},
	}); err != nil {
		t.Fatal(err)
	}

	s = "'hello world'"
	tokens, err = Tokenizer(s)

	if err := testCheckTokens(tokens, []Token{
		Token{TK_STRING, "'hello world'"},
	}); err != nil {
		t.Fatal(err)
	}
}

func TestLexer_StringMix(t *testing.T) {
	s := `"""hello world"`

	tokens, err := Tokenizer(s)

	if err != nil {
		t.Fatal(err)
	}

	expected := []Token{
		Token{TK_STRING, `"""hello world"`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `"'foo'", "\"foo", "'foo", '"foo', '\'foo', '''foo'`

	tokens, err = Tokenizer(s)

	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_STRING, `"'foo'"`},
		Token{TK_COMMA, `,`},
		Token{TK_STRING, `"\"foo"`},
		Token{TK_COMMA, `,`},
		Token{TK_STRING, `"'foo"`},
		Token{TK_COMMA, `,`},
		Token{TK_STRING, `'"foo'`},
		Token{TK_COMMA, `,`},
		Token{TK_STRING, `'\'foo'`},
		Token{TK_COMMA, `,`},
		Token{TK_STRING, `'''foo'`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `"abc`

	tokens, err = Tokenizer(s)
	if err == nil {
		t.Fatal("expect error")
	}
}

func TestLexer_Number(t *testing.T) {
	s := `12345`

	tokens, err := Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected := []Token{
		Token{TK_INTEGER, `12345`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `0x12345`

	tokens, err = Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_INTEGER, `0x12345`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `0b0010`

	tokens, err = Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_INTEGER, `0b0010`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}
}

func TestLexer_NumberFloat(t *testing.T) {
	s := `3.14`

	tokens, err := Tokenizer(s)

	if err != nil {
		t.Fatal(err)
	}

	expected := []Token{
		Token{TK_FLOAT, `3.14`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `.14`

	tokens, err = Tokenizer(s)

	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_FLOAT, `.14`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `3.`
	tokens, err = Tokenizer(s)

	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_FLOAT, `3.`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `3e+10`
	tokens, err = Tokenizer(s)

	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_FLOAT, `3e+10`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `3.1e+10`
	tokens, err = Tokenizer(s)

	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_FLOAT, `3.1e+10`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

}

func TestLexer_Comment(t *testing.T) {
	s := `-- abc`

	tokens, err := Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected := []Token{
		Token{TK_COMMENT, `-- abc`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `/* abc */`

	tokens, err = Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_COMMENT, `/* abc */`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `/*! abc */`
	tokens, err = Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_COMMENT_MYSQL, `/*! abc */`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}
}

func TestTexer_CommentMix(t *testing.T) {
	s := `--1`

	tokens, err := Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected := []Token{
		Token{TK_MINUS, `-`},
		Token{TK_MINUS, `-`},
		Token{TK_INTEGER, `1`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = "-- abc\n123"

	tokens, err = Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_COMMENT, `-- abc`},
		Token{TK_INTEGER, `123`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `123/*abc*/123`

	tokens, err = Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_INTEGER, `123`},
		Token{TK_COMMENT, `/*abc*/`},
		Token{TK_INTEGER, `123`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = "123 /*abc"
	tokens, err = Tokenizer(s)
	if err == nil {
		t.Fatal("expect error")
	}
}

func TestLexer_Variable(t *testing.T) {
	s := `select @@abc`

	tokens, err := Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected := []Token{
		Token{TK_SQL_SELECT, `select`},
		Token{TK_SYS_VARIABLE, `@@abc`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `select @abc`

	tokens, err = Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_SQL_SELECT, `select`},
		Token{TK_VARIABLE, `@abc`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `select @abc.abc`

	tokens, err = Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_SQL_SELECT, `select`},
		Token{TK_VARIABLE, `@abc.abc`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

}

func TestLexer_Identifier(t *testing.T) {
	s := `select * from abc`

	tokens, err := Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected := []Token{
		Token{TK_SQL_SELECT, `select`},
		Token{TK_STAR, `*`},
		Token{TK_SQL_FROM, `from`},
		Token{TK_LITERAL, `abc`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `1e+1e`

	tokens, err = Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_FLOAT, `1e+1`},
		Token{TK_LITERAL, `e`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `abc.123`

	tokens, err = Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_LITERAL, `abc`},
		Token{TK_DOT, `.`},
		Token{TK_LITERAL, `123`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `select abc.123, abc()`

	tokens, err = Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_SQL_SELECT, `select`},
		Token{TK_LITERAL, `abc`},
		Token{TK_DOT, `.`},
		Token{TK_LITERAL, `123`},
		Token{TK_COMMA, `,`},
		Token{TK_FUNCTION, `abc`},
		Token{TK_LPAREN, `(`},
		Token{TK_RPAREN, `)`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

	s = `e.1e+10`

	tokens, err = Tokenizer(s)
	if err != nil {
		t.Fatal(err)
	}

	expected = []Token{
		Token{TK_LITERAL, `e`},
		Token{TK_DOT, `.`},
		Token{TK_LITERAL, `1e`},
		Token{TK_PLUS, `+`},
		Token{TK_INTEGER, `10`},
	}

	if err := testCheckTokens(tokens, expected); err != nil {
		t.Fatal(err)
	}

}
