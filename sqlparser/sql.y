%{
package sqlparser   

func SetParseTree(yylex interface{}, stmt Statement) {
    tn := yylex.(*Tokenizer)
    tn.ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
    tn := yylex.(*Tokenizer)
    tn.AllowComments = allow
}

func ForceEOF(yylex interface{}) {
    tn := yylex.(*Tokenizer)
    tn.ForceEOF = true
}

%}

%union {
    node *Node
    stmt Stmt
    expr Expr
    num  int
}

%token <node> SELECT INSERT UPDATE DELETE REPLACE FROM WHERE GROUP HAVING ORDER BY LIMIT COMMENT FOR
%token <node> ALL DISTINCT AS EXISTS IN IS LIKE BETWEEN NULL ASC DESC VALUES INTO DUPLICATE KEY DEFAULT SET LOCK
%token <node> BEGIN COMMIT ROLLBACK
%token <node> ID STRING NUMBER NAME
%token <node> LE GE NE NULL_SAFE_EQUAL
%token <node> LEX_ERROR
%token <node> '(' '=' '<' '>' '~'

%token <node> LOW_PRIORITY QUICK IGNORE


%left <node> UNION MINUS EXCEPT INTERSECT
%left <node> ','
%left <node> JOIN STRAIGHT_JOIN LEFT RIGHT INNER OUTER CROSS NATURAL USE FORCE
%left <node> ON
%left <node> AND OR
%right <node> NOT
%left <node> '&' '|' '^'
%left <node> '+' '-'
%left <node> '*' '/' '%'
%nonassoc <node> '.'
%left <node> UNARY
%right <node> CASE, WHEN, THEN, ELSE
%left <node> END



%start any_command

%type <stmt> command
%type <stmt> select_statement update_statement delete_statement insert_statement replace_statement set_statement
%type <num> delete_opts

%%

any_command: 
    command 
    {
        SetParseTree(yylex, $1)
    }

command:
    select_statement
|   update_statement
|   delete_statement
|   insert_statement
|   replace_statement
|   set_statement 

select_statement:
    SELECT
    {

    }

update_statement:
    UPDATE
    {

    }

delete_statement:
    DELETE delete_opts FROM NAME
    {
        $$ = &DeleteStmt{Opts: $2, Table: $4}
    }

insert_statement:
    INSERT
    {
    
    }

replace_statement:
    REPLACE
    {

    }

set_statement: 
    {

    }


delete_opts: 
    delete_opts LOW_PRIORITY 
    {
        $$ = $1 | 1
    }
|   delete_opts QUICK 
    {    
        $$ = $1 | 2
    }
|   delete_opts IGNORE 
    {
        $$ = $1 | 4
    }
|   {   
        $$ = 0 
    }
