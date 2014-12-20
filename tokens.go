package main

type Token int

const (
	tInvalid Token = iota

	tLiteral

	tAND
	tARRAY
	tASSIGN
	tBEGIN
	tBOOLEAN
	tBREAK
	tCHAR
	tCONTINUE
	tDIV
	tDO
	tDOTDOT
	tDOWNTO
	tELSE
	tEND
	tFALSE
	tFNUMBER
	tFOR
	tFUNCTION
	tGE
	tGOTO
	tID
	tIF
	tINTEGER
	tLE
	tMOD
	tNE
	tNOT
	tNUMBER
	tOF
	tOR
	tPROCEDURE
	tPROGRAM
	tQSTRING
	tREAL
	tRECORD
	tREPEAT
	tSTRING
	tTHEN
	tTHIS
	tTO
	tTRUE
	tTYPE
	tUNTIL
	tVAR
	tWHILE
)

var reserved = map[string]Token{
	"and":       tAND,
	"array":     tARRAY,
	"begin":     tBEGIN,
	"boolean":   tBOOLEAN,
	"break":     tBREAK,
	"char":      tCHAR,
	"continue":  tCONTINUE,
	"div":       tDIV,
	"do":        tDO,
	"downto":    tDOWNTO,
	"else":      tELSE,
	"end":       tEND,
	"false":     tFALSE,
	"for":       tFOR,
	"function":  tFUNCTION,
	"goto":      tGOTO,
	"if":        tIF,
	"integer":   tINTEGER,
	"mod":       tMOD,
	"not":       tNOT,
	"of":        tOF,
	"or":        tOR,
	"procedure": tPROCEDURE,
	"program":   tPROGRAM,
	"real":      tREAL,
	"record":    tRECORD,
	"repeat":    tREPEAT,
	"string":    tSTRING,
	"then":      tTHEN,
	"to":        tTO,
	"true":      tTRUE,
	"type":      tTYPE,
	"until":     tUNTIL,
	"var":       tVAR,
	"while":     tWHILE,
}

type Lexeme struct {
	T Token
	S string
	L int
}
