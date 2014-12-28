package main

import (
	"fmt"
	"bytes"
	"strconv"
)

%% machine scanner;
%% write data;

type tok struct {
    t int
    yy yySymType
}

func lex(data []byte) []tok {

	cs, p, pe, eof := 0, 0, len(data), len(data)
	ts, te, act := 0, 0, 0

	_, _, _ = ts, te, act

	lineno := 1

	var tokens []tok

	add := func(t int) {
	    tokens = append(tokens, tok{t:t})
	}

	addstr := func(t int, s string) {
	    tokens = append(tokens, tok{t:t, yy:yySymType{s:s}})
	}

	addint := func(t int, i int) {
	    tokens = append(tokens, tok{t:t, yy:yySymType{i:i}})
	}

	addf := func(t int, f float64) {
	    tokens = append(tokens, tok{t:t, yy:yySymType{f:f}})
	}

	%%{

	    main := |*
		[\-@^:;,=><[\](){}+/.*] => { add(int(data[ts])) };
		':=' => { add(tASSIGN) };
		'\.\.' => { add(tDOTDOT) };
		'>=' => { add(tGE) };
		'<=' => { add(tLE) };
		'<>' => { add(tNE) };
		'\'' [^']* '\''  => { addstr(tQSTRING, string(data[ts:te])) };
		'{' [^}]* '}'  => { lineno += bytes.Count(data[ts:te], []byte{'\n'}) };
		[ \t] => { };
		'\n' => { lineno++ };
		digit+ '.' digit+ {
		    f, _ := strconv.ParseFloat(string(data[ts:te]), 64)
		    addf(tFNUMBER, f)
		};
		digit+ {
		    i, _ := strconv.Atoi(string(data[ts:te]))
		    addint(tNUMBER, i)
		};
		[A-Za-z_][A-Za-z0-9_]* {
			t, ok := reserved[string(data[ts:te])];
			if !ok {
				t = tID
			};
			addstr(t, string(data[ts:te]))
		};
	    *|;

	    write init;
	    write exec;
	}%%

	return tokens
}

type pascalLexer []tok

func (p *pascalLexer) Lex(lval *yySymType) int  {
    if len(*p) == 0 {
	    return 0
    }
    t := (*p)[0]
    *p = (*p)[1:]
    *lval = t.yy
    return t.t
}

func (p *pascalLexer) Error(s string) {
    fmt.Println("syntax error:", s)
}

