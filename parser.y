%{

package main

%}

%union{
        f float64
        i int
        s string
}

%token <f> tFNUMBER
%token <i> tNUMBER
%token <s> tID tQSTRING

%token	tAND tARRAY tASSIGN tBEGIN tBOOLEAN tBREAK tCHAR tCONTINUE tDIV tDO tDOTDOT
%token	tDOWNTO tELSE tEND tFALSE tFOR tFUNCTION tGE tGOTO tIF tINTEGER tLE tMOD
%token	tNE tNOT tOF tOR tPROCEDURE tPROGRAM tREAL tRECORD tREPEAT tSTRING
%token	tTHEN tTO tTRUE tTYPE tUNTIL tVAR tWHILE

%nonassoc '<' '>' '=' tLE tGE tNE
%left '+'  '-' tOR
%left '*'  '/' tAND tDIV tMOD
%right UMINUS tNOT

%%

pascal_program : tPROGRAM tID ';' decls subprog_decls compound_stmt '.' ;

id_list : id_list ',' tID
        | tID
        ;

decls: decls tVAR id_list ':' ptype ';'
     | decls tTYPE tID '=' ptype ';'
     |  /* empty */
     ;

ptype: standard_type
     | tID
     | tARRAY '[' tNUMBER tDOTDOT tNUMBER ']' tOF standard_type
     | '^' ptype
     | tRECORD param_list_list ';' tEND
     | tFUNCTION arguments ':' standard_type
     | tPROCEDURE arguments
     ;

standard_type: tINTEGER
             | tREAL
             | tCHAR
             | tBOOLEAN
             | tSTRING
             ;

subprog_decls: subprog_decls subprog_decl
             | /* empty */
             ;

subprog_decl : subprog_head decls compound_stmt ';'
              ;

subprog_head: tFUNCTION tID arguments ':' standard_type ';'
            | tPROCEDURE tID arguments ';'
            ;

arguments: '(' param_list_list ')'
         | /* empty */
         ;

param_list_list: param_list_list ';' param_list
               | param_list
               ;

param_list: tVAR id_list ':' ptype
          | id_list ':' ptype
          ;

compound_stmt : tBEGIN opt_stmts tEND

opt_stmts : stmt_list
          | /* empty */
          ;

stmt_list : stmt_list stmt ';'
          | stmt ';'
          ;

stmt : variable tASSIGN expr
     | tBREAK
     | tCONTINUE
     | tID '(' expr_list ')'
     | tID
     | compound_stmt
     | tIF expr tTHEN stmt
     | tIF expr tTHEN stmt tELSE stmt
     | tFOR tID tASSIGN expr tTO expr tDO stmt
     | tWHILE expr tDO stmt
     | tREPEAT stmt_list tUNTIL expr
     ;

expr_list: expr_list ',' expr
         | expr
         ;

expr: tID '(' expr_list ')' 
    | '@' variable
    | '+' expr %prec UMINUS
    | '-' expr %prec UMINUS
    | variable
    | expr tAND expr
    | expr tDIV expr
    | expr '<' expr
    | expr '=' expr
    | expr '>' expr
    | expr '-' expr
    | expr '/' expr
    | expr '*' expr
    | expr '+' expr
    | expr tGE expr
    | expr tLE expr
    | expr tMOD expr
    | expr tNE expr
    | expr tOR expr
    | '(' expr ')'
    | tNOT expr
    | tNUMBER
    | tFNUMBER
    | tTRUE
    | tFALSE
    | tQSTRING
    ;

variable: tID
    | variable '[' expr ']'
    | variable '.' tID
    | variable '^'
    ;
