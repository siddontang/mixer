%{
package sqlparser   

func SetParseTree(l interface{}, stmt Statement) {
    lexer := l.(*Lexer)
    lexer.ParseTree = stmt
}
%}

%union {
    node *Node
    statement Statement
}

%token <node> SELECT
%token <statement> command
%token <statement> select_statement

%start any_command

%%

any_command: 
    command 
    {
        SetParseTree(yylex, $1)
    }

command
