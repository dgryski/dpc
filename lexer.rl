package main

import (
	"bytes"
)

%% machine scanner;
%% write data;

func lex(data []byte) []Lexeme {

	cs, p, pe, eof := 0, 0, len(data), len(data)
	ts, te, act := 0, 0, 0

	_, _, _ = ts, te, act

	lineno := 1

	var tokens []Lexeme

	add := func(t Token) {
		tokens = append(tokens, Lexeme{t, string(data[ts:te]), lineno})
	}

	%%{

	    main := |*
		[\-@^:;,=><[\](){}+/.*] => { add(tLiteral) };
		':=' => { add(tASSIGN) };
		'\.\.' => { add(tDOTDOT) };
		'>=' => { add(tGE) };
		'<=' => { add(tLE) };
		'<>' => { add(tNE) };
		'\'' [^']* '\''  => { add(tQSTRING) };
		'{' [^}]* '}'  => { lineno += bytes.Count(data[ts:te], []byte{'\n'}) };
		[ \t] => { };
		'\n' => { lineno++ };
		digit+ '.' digit+ { add(tFNUMBER) };
		digit+ { add(tNUMBER) };
		[A-Za-z_][A-Za-z0-9_]* {
			t, ok := reserved[string(data[ts:te])];
			if !ok {
				t = tID
			};
			add(t)
		};
	    *|;

	    write init;
	    write exec;
	}%%

	return tokens
}
