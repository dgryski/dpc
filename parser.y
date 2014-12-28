%{

package main

import "log"

%}

%union{
        f float64
        i int
        s string

        vars []pVar
        typedefs []typTypedef
        strings []string
        ptyp pType
        program pProgram
        functions []pFunction
        function pFunction
        decls pDecls
        expr expr
        exprs []expr
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

%type <strings> id_list
%type <ptyp> standard_type ptype
%type <vars> param_list param_list_list arguments
%type <decls> decls
%type <functions> subprog_decls
%type <function> subprog_decl subprog_head
%type <exprs> expr_list
%type <expr> expr variable
%%

pascal_program : tPROGRAM tID ';' decls subprog_decls compound_stmt '.' { log.Printf("%#v", pProgram{name:$2, vars:$4.vars, types:$4.types, subprogs:$5}) } ;

id_list : id_list ',' tID { $$ = append($1, $3); }
        | tID { $$ = append($$, $1) }
        ;

decls: decls tVAR id_list ':' ptype ';' { for _, id := range $3 { $$.vars = append($$.vars, pVar{name:id, typ:$5}) } }
     | decls tTYPE tID '=' ptype ';' { $$.types = append($1.types, typTypedef{name:$3, typ:$5}) }
     | { $$ = pDecls{} } /* empty */
     ;

ptype: standard_type { $$ = $1 }
     | tID { $$ = typTypedef{name:$1} }
     | tARRAY '[' tNUMBER tDOTDOT tNUMBER ']' tOF standard_type { $$ = typArray{start:$3, end:$5, typ:$8} }
     | '^' ptype { $$ = typPointer{typ:$2} }
     | tRECORD param_list_list ';' tEND { $$ = typRecord{fields:$2} }
     | tFUNCTION arguments ':' standard_type { $$ = typVoid{} }
     | tPROCEDURE arguments { $$ = typVoid{} }
     ;

standard_type: tINTEGER { $$ = typPrimitive{primInt} }
             | tREAL { $$ = typPrimitive{primReal} }
             | tCHAR { $$ = typPrimitive{primChar} }
             | tBOOLEAN { $$ = typPrimitive{primBool} }
             | tSTRING { $$ = typPrimitive{primString} }
             ;

subprog_decls: subprog_decls subprog_decl { $$ = append($1, $2) }
             | { $$ = nil }/* empty */
             ;

subprog_decl : subprog_head decls compound_stmt ';' { $1.decls = $2.vars; $$ = $1 }
              ;

subprog_head: tFUNCTION tID arguments ':' standard_type ';' { $$ = pFunction{name:$2, args:$3, ret:$5} }
            | tPROCEDURE tID arguments ';' { $$ = pFunction{name:$2, args:$3, ret:typVoid{}} }
            ;

arguments: '(' param_list_list ')' { $$ = $2 }
         | /* empty */ { $$ = nil }
         ;

param_list_list: param_list_list ';' param_list { $$ = append($1, $3...) }
               | param_list { $$ = $1 }
               ;

param_list: tVAR id_list ':' ptype  { for _, id := range $2 { $$ = append($$, pVar{name:id, typ:$4}) } }
          | id_list ':' ptype { for _, id := range $1 { $$ = append($$, pVar{name:id, typ:$3}) } }
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

expr_list: expr_list ',' expr { $$ = append($1, $3) }
         | expr { $$ = []expr{$1} }
         ;

expr: tID '(' expr_list ')' { $$ = expCall{fn:pVar{name:$1}, args: $3} }
    | '@' variable { $$ = expUnop{op:unopAt, e:$2} }
    | '+' expr %prec UMINUS { $$ = expUnop{op:unopPlus, e:$2} }
    | '-' expr %prec UMINUS { $$ = expUnop{op:unopMinus,  e:$2} }
    | variable { $$ = $1 }
    | expr tAND expr { $$ = expBinop{op:binAND, left:$1, right: $3} }
    | expr tDIV expr { $$ = expBinop{op:binDIV, left:$1, right: $3} }
    | expr '<' expr { $$ = expBinop{op:binLT, left:$1, right: $3} }
    | expr '=' expr { $$ = expBinop{op:binEQ, left:$1, right: $3} }
    | expr '>' expr { $$ = expBinop{op:binGT, left:$1, right: $3} }
    | expr '-' expr { $$ = expBinop{op:binSUB, left:$1, right: $3} }
    | expr '/' expr { $$ = expBinop{op:binFDIV, left:$1, right: $3} }
    | expr '*' expr { $$ = expBinop{op:binMUL, left:$1, right: $3} }
    | expr '+' expr { $$ = expBinop{op:binADD, left:$1, right: $3} }
    | expr tGE expr { $$ = expBinop{op:binGE, left:$1, right: $3} }
    | expr tLE expr { $$ = expBinop{op:binLE, left:$1, right: $3} }
    | expr tMOD expr { $$ = expBinop{op:binMOD, left:$1, right: $3} }
    | expr tNE expr { $$ = expBinop{op:binNE, left:$1, right: $3} }
    | expr tOR expr { $$ = expBinop{op:binOR, left:$1, right: $3} }
    | '(' expr ')' { $$ = $2 }
    | tNOT expr { $$ = expUnop{op:unopNot, e:$2} }
    | tNUMBER   { $$ = expConst{t:primInt, i:$1} }
    | tFNUMBER{ $$ = expConst{t:primInt, f:$1} }
    | tTRUE { $$ = expConst{t:primInt, b:true} }
    | tFALSE { $$ = expConst{t:primInt, b:false} }
    | tQSTRING{ $$ = expConst{t:primInt, s:$1} }

    ;

variable: tID { $$ = expId{name:$1} }
    | variable '[' expr ']' { $$ = expBinop{op:binArrayIndex, left:$1, right:$3} }
    | variable '.' tID { $$ = expField{e:$1, field:pVar{name:$3}} }
    | variable '^' { $$ = expUnop{op:unopPtr, e:$1} }
    ;
