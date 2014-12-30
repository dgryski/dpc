%{

package main

var program varProgram
%}

%union{
        f float64
        i int
        s string

        vars []varId
        typedefs []typTypedef
        strings []string
        ptyp pType
        program varProgram
        functions []varFunction
        function varFunction
        decls pDecls
        expr expr
        exprs []expr
        stmt stmt
        stmts []stmt
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
%type <stmt> stmt compound_stmt
%type <stmts> stmt_list opt_stmts
%%

pascal_program : tPROGRAM tID ';' decls subprog_decls compound_stmt '.' { program = varProgram{name:$2, vars:$4.vars, types:$4.types, subprogs:$5, body:$6} } ;

id_list : id_list ',' tID { $$ = append($1, $3); }
        | tID { $$ = append($$, $1) }
        ;

decls: decls tVAR id_list ':' ptype ';' { for _, id := range $3 { $$.vars = append($$.vars, varId{name:id, typ:$5}) } }
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

subprog_decl : subprog_head decls compound_stmt ';' { $1.decls = $2.vars; $1.body = $3; $$ = $1 }
              ;

subprog_head: tFUNCTION tID arguments ':' standard_type ';' { $$ = varFunction{name:$2, args:$3, ret:varId{typ:$5}}}
            | tPROCEDURE tID arguments ';' { $$ = varFunction{name:$2, args:$3, ret:varId{typ:typVoid{}}} }
            ;

arguments: '(' param_list_list ')' { $$ = $2 }
         | /* empty */ { $$ = nil }
         ;

param_list_list: param_list_list ';' param_list { $$ = append($1, $3...) }
               | param_list { $$ = $1 }
               ;

param_list: tVAR id_list ':' ptype  { for _, id := range $2 { $$ = append($$, varId{name:id, typ:$4}) } }
          | id_list ':' ptype { for _, id := range $1 { $$ = append($$, varId{name:id, typ:$3}) } }
          ;

compound_stmt : tBEGIN opt_stmts tEND { $$ = stmBlock{stmts:$2} }

opt_stmts : stmt_list { $$ = $1 }
          | /* empty */ { $$ = nil }
          ;

stmt_list : stmt_list stmt ';' { $$ = append($1, $2) }
          | stmt ';' { $$ = []stmt{$1} }
          ;

stmt : variable tASSIGN expr { $$ = stmAssign{id:$1, e:$3} }
     | tBREAK { $$ = stmBreak{} }
     | tCONTINUE { $$ = stmContinue{} }
     | tID '(' expr_list ')' { $$ = stmCall{fn:varId{name:$1}, args:$3} }
     | tID  { $$ = stmCall{fn:varId{name:$1}} }
     | compound_stmt { $$ = $1 }
     | tIF expr tTHEN stmt { $$ = stmIf{cond:$2, ifTrue:$4} }
     | tIF expr tTHEN stmt tELSE stmt { $$ = stmIf{cond:$2, ifTrue:$4, ifFalse:$6} }
     | tFOR tID tASSIGN expr tTO expr tDO stmt { $$ = stmFor{counter:expId{name:$2}, expr1:$4, expr2:$6, body:$8} }
     | tWHILE expr tDO stmt { $$ = stmWhile{e:$2, body:$4} }
     | tREPEAT stmt_list tUNTIL expr { $$ = stmRepeat{e:$4, body:stmBlock{$2}} }
     ;

expr_list: expr_list ',' expr { $$ = append($1, $3) }
         | expr { $$ = []expr{$1} }
         ;

expr: tID '(' expr_list ')' { $$ = expCall{fn:varId{name:$1}, args: $3} }
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
    | tNUMBER { $$ = expConst{t:primInt, i:$1} }
    | tFNUMBER { $$ = expConst{t:primReal, f:$1} }
    | tTRUE { $$ = expConst{t:primBool, b:true} }
    | tFALSE { $$ = expConst{t:primBool, b:false} }
    | tQSTRING { $$ = expConst{t:primString, s:$1}; }
    ;

variable: tID { $$ = expId{name:$1} }
    | variable '[' expr ']' { $$ = expBinop{op:binArrayIndex, left:$1, right:$3} }
    | variable '.' tID { $$ = expField{e:$1, field:varId{name:$3}} }
    | variable '^' { $$ = expUnop{op:unopPtr, e:$1} }
    ;
