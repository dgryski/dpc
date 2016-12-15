//line parser.y:2
package main

import __yyfmt__ "fmt"

//line parser.y:3
var program varProgram

//line parser.y:8
type yySymType struct {
	yys int
	f   float64
	i   int
	s   string

	vars      []*varId
	typedefs  []*typTypedef
	strings   []string
	ptyp      pType
	program   varProgram
	functions []*varFunction
	function  *varFunction
	decls     pDecls
	expr      expr
	exprs     []expr
	stmt      stmt
	stmts     []stmt
}

const tFNUMBER = 57346
const tNUMBER = 57347
const tID = 57348
const tQSTRING = 57349
const tAND = 57350
const tARRAY = 57351
const tASSIGN = 57352
const tBEGIN = 57353
const tBOOLEAN = 57354
const tBREAK = 57355
const tCHAR = 57356
const tCONTINUE = 57357
const tDIV = 57358
const tDO = 57359
const tDOTDOT = 57360
const tDOWNTO = 57361
const tELSE = 57362
const tEND = 57363
const tFALSE = 57364
const tFOR = 57365
const tFUNCTION = 57366
const tGE = 57367
const tGOTO = 57368
const tIF = 57369
const tINTEGER = 57370
const tLE = 57371
const tMOD = 57372
const tNE = 57373
const tNOT = 57374
const tOF = 57375
const tOR = 57376
const tPROCEDURE = 57377
const tPROGRAM = 57378
const tREAL = 57379
const tRECORD = 57380
const tREPEAT = 57381
const tSTRING = 57382
const tTHEN = 57383
const tTO = 57384
const tTRUE = 57385
const tTYPE = 57386
const tUNTIL = 57387
const tVAR = 57388
const tWHILE = 57389
const UMINUS = 57390

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"tFNUMBER",
	"tNUMBER",
	"tID",
	"tQSTRING",
	"tAND",
	"tARRAY",
	"tASSIGN",
	"tBEGIN",
	"tBOOLEAN",
	"tBREAK",
	"tCHAR",
	"tCONTINUE",
	"tDIV",
	"tDO",
	"tDOTDOT",
	"tDOWNTO",
	"tELSE",
	"tEND",
	"tFALSE",
	"tFOR",
	"tFUNCTION",
	"tGE",
	"tGOTO",
	"tIF",
	"tINTEGER",
	"tLE",
	"tMOD",
	"tNE",
	"tNOT",
	"tOF",
	"tOR",
	"tPROCEDURE",
	"tPROGRAM",
	"tREAL",
	"tRECORD",
	"tREPEAT",
	"tSTRING",
	"tTHEN",
	"tTO",
	"tTRUE",
	"tTYPE",
	"tUNTIL",
	"tVAR",
	"tWHILE",
	"'<'",
	"'>'",
	"'='",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"UMINUS",
	"';'",
	"'.'",
	"','",
	"':'",
	"'['",
	"']'",
	"'^'",
	"'('",
	"')'",
	"'@'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 25,
	20, 39,
	56, 39,
	-2, 74,
	-1, 131,
	25, 0,
	29, 0,
	31, 0,
	48, 0,
	49, 0,
	50, 0,
	-2, 55,
	-1, 132,
	25, 0,
	29, 0,
	31, 0,
	48, 0,
	49, 0,
	50, 0,
	-2, 56,
	-1, 133,
	25, 0,
	29, 0,
	31, 0,
	48, 0,
	49, 0,
	50, 0,
	-2, 57,
	-1, 138,
	25, 0,
	29, 0,
	31, 0,
	48, 0,
	49, 0,
	50, 0,
	-2, 62,
	-1, 139,
	25, 0,
	29, 0,
	31, 0,
	48, 0,
	49, 0,
	50, 0,
	-2, 63,
	-1, 141,
	25, 0,
	29, 0,
	31, 0,
	48, 0,
	49, 0,
	50, 0,
	-2, 65,
}

const yyNprod = 78
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 444

var yyAct = [...]int{

	85, 67, 66, 114, 116, 84, 62, 127, 127, 113,
	150, 63, 15, 158, 126, 50, 21, 40, 149, 42,
	101, 44, 41, 119, 43, 172, 68, 22, 45, 69,
	59, 155, 77, 112, 76, 18, 22, 38, 160, 79,
	64, 81, 82, 154, 72, 124, 22, 118, 74, 104,
	105, 117, 106, 107, 111, 73, 80, 75, 71, 2,
	78, 34, 162, 102, 42, 39, 4, 41, 36, 43,
	34, 152, 16, 120, 34, 35, 22, 38, 174, 122,
	123, 121, 70, 11, 16, 37, 164, 165, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	141, 142, 22, 128, 8, 87, 7, 143, 26, 145,
	157, 147, 115, 88, 148, 9, 8, 11, 7, 77,
	151, 76, 20, 5, 115, 22, 146, 98, 156, 108,
	13, 54, 53, 46, 57, 74, 31, 16, 103, 83,
	61, 14, 65, 58, 75, 33, 32, 78, 17, 56,
	94, 93, 3, 60, 161, 163, 170, 166, 161, 52,
	168, 153, 1, 87, 19, 169, 12, 10, 6, 0,
	55, 88, 0, 22, 167, 0, 175, 0, 48, 49,
	96, 0, 0, 0, 97, 98, 99, 22, 173, 100,
	51, 0, 47, 0, 87, 0, 0, 0, 0, 0,
	0, 0, 88, 89, 91, 90, 95, 92, 94, 93,
	0, 96, 0, 0, 0, 97, 98, 99, 0, 144,
	100, 0, 87, 0, 0, 0, 0, 0, 0, 0,
	88, 0, 0, 87, 89, 91, 90, 95, 92, 94,
	93, 88, 171, 0, 98, 0, 0, 125, 100, 0,
	96, 0, 0, 0, 97, 98, 99, 0, 0, 100,
	0, 0, 0, 87, 0, 95, 92, 94, 93, 0,
	0, 88, 0, 89, 91, 90, 95, 92, 94, 93,
	96, 0, 0, 0, 97, 98, 99, 0, 0, 100,
	0, 0, 0, 87, 0, 0, 0, 159, 0, 0,
	0, 88, 109, 89, 91, 90, 95, 92, 94, 93,
	96, 0, 0, 0, 97, 98, 99, 0, 0, 100,
	0, 0, 0, 87, 0, 0, 0, 0, 0, 0,
	0, 88, 0, 89, 91, 90, 95, 92, 94, 93,
	96, 0, 0, 0, 97, 98, 99, 0, 0, 100,
	0, 0, 0, 87, 0, 0, 86, 0, 0, 0,
	0, 88, 0, 89, 91, 90, 95, 92, 94, 93,
	96, 0, 0, 0, 97, 98, 99, 0, 0, 100,
	0, 0, 0, 25, 0, 0, 0, 0, 11, 0,
	23, 0, 24, 89, 91, 90, 95, 92, 94, 93,
	28, 0, 25, 0, 27, 0, 0, 11, 0, 23,
	0, 24, 0, 0, 0, 0, 30, 0, 0, 28,
	0, 0, 110, 27, 29, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 30, 0, 0, 0, 0,
	0, 0, 0, 29,
}
var yyPact = [...]int{

	23, -1000, 146, 10, -1000, 60, 106, 131, 142, -22,
	-1000, 396, -1000, 140, 139, 16, -1000, 18, -1000, 64,
	396, 9, 7, -1000, -1000, -42, -1000, 127, 137, 127,
	396, 72, -52, -52, 136, 20, 20, -1000, 0, -1000,
	127, 127, 133, -1000, 127, 315, -43, 132, 127, 127,
	-38, 127, 127, -1000, -1000, -1000, -1000, -1000, 119, 285,
	377, -2, -26, 78, -5, -1000, -9, -1000, -1000, -37,
	20, 78, -52, -52, -1000, -1000, -1000, -1000, -1000, -11,
	-1000, 345, 186, -1000, -50, 345, 396, 127, 127, 127,
	127, 127, 127, 127, 127, 127, 127, 127, 127, 127,
	127, 127, -38, -1000, -1000, -1000, 155, -1000, 127, 396,
	127, -1000, 107, -46, -1000, 131, 12, -1000, -1000, 156,
	-1000, -13, -28, -1000, -1000, -1000, -1000, 127, 90, -1000,
	-1000, 214, 214, 214, 97, -1000, -1000, 97, 214, 214,
	-1000, 214, 97, -51, -1000, 255, -1000, 345, -18, -1000,
	78, 3, 20, 68, 66, 107, 345, 396, -1000, 127,
	-1000, -1000, 20, -1000, 151, -1000, -1000, -1000, 225, -1000,
	-36, 396, 45, -1000, 107, -1000,
}
var yyPgo = [...]int{

	0, 4, 1, 2, 3, 9, 6, 123, 168, 167,
	166, 5, 0, 15, 16, 108, 122, 164, 162,
}
var yyR1 = [...]int{

	0, 18, 1, 1, 7, 7, 7, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 2, 2, 2, 8,
	8, 9, 10, 10, 6, 6, 5, 5, 4, 4,
	15, 17, 17, 16, 16, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 11, 11, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 13, 13, 13, 13,
}
var yyR2 = [...]int{

	0, 7, 3, 1, 6, 6, 0, 1, 1, 8,
	2, 4, 4, 2, 1, 1, 1, 1, 1, 2,
	0, 4, 6, 4, 3, 0, 3, 1, 4, 3,
	3, 1, 0, 3, 2, 3, 1, 1, 4, 1,
	1, 4, 6, 8, 4, 4, 3, 1, 4, 2,
	2, 2, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 1,
	1, 1, 1, 1, 1, 4, 3, 2,
}
var yyChk = [...]int{

	-1000, -18, 36, 6, 56, -7, -8, 46, 44, -15,
	-9, 11, -10, 24, 35, -1, 6, 6, 57, -17,
	-16, -14, -13, 13, 15, 6, -15, 27, 23, 47,
	39, -7, 6, 6, 58, 59, 50, 21, -14, 56,
	10, 60, 57, 62, 63, -12, 6, 65, 51, 52,
	-13, 63, 32, 5, 4, 43, 22, 7, 6, -12,
	-16, -15, -6, 63, -6, 6, -3, -2, 6, 9,
	62, 38, 24, 35, 28, 37, 14, 12, 40, -3,
	56, -12, -12, 6, -11, -12, 41, 8, 16, 48,
	50, 49, 52, 54, 53, 51, 25, 29, 30, 31,
	34, 63, -13, 6, -12, -12, -12, -12, 10, 17,
	45, 56, 59, -5, -4, 46, -1, 56, 56, 60,
	-3, -5, -6, -6, 56, 61, 64, 58, -14, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -11, 64, -12, -14, -12, -2, 64,
	56, -1, 59, 5, 56, 59, -12, 20, 64, 42,
	56, -4, 59, -3, 18, 21, -2, -14, -12, -3,
	5, 17, 61, -14, 33, -2,
}
var yyDef = [...]int{

	0, -2, 0, 0, 6, 20, 0, 0, 0, 0,
	19, 32, 6, 0, 0, 0, 3, 0, 1, 0,
	31, 0, 0, 36, 37, -2, 40, 0, 0, 0,
	0, 0, 25, 25, 0, 0, 0, 30, 0, 34,
	0, 0, 0, 77, 0, 0, 74, 0, 0, 0,
	52, 0, 0, 69, 70, 71, 72, 73, 0, 0,
	0, 0, 0, 0, 0, 2, 0, 7, 8, 0,
	0, 0, 25, 25, 14, 15, 16, 17, 18, 0,
	33, 35, 0, 76, 0, 47, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 49, 74, 50, 51, 0, 68, 0, 0,
	0, 21, 0, 0, 27, 0, 0, 23, 4, 0,
	10, 0, 0, 13, 5, 75, 38, 0, 41, 53,
	54, -2, -2, -2, 58, 59, 60, 61, -2, -2,
	64, -2, 66, 0, 67, 0, 44, 45, 0, 24,
	0, 0, 0, 0, 0, 0, 46, 0, 48, 0,
	22, 26, 0, 29, 0, 11, 12, 42, 0, 28,
	0, 0, 0, 43, 0, 9,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	63, 64, 53, 51, 58, 52, 57, 54, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 59, 56,
	48, 50, 49, 3, 65, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 60, 3, 61, 62,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 55,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:53
		{
			program = varProgram{varDecl: varDecl{name: yyDollar[2].s}, vars: yyDollar[4].decls.vars, types: yyDollar[4].decls.types, subprogs: yyDollar[5].functions, body: yyDollar[6].stmt}
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:55
		{
			yyVAL.strings = append(yyDollar[1].strings, yyDollar[3].s)
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:56
		{
			yyVAL.strings = append(yyVAL.strings, yyDollar[1].s)
		}
	case 4:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:59
		{
			for _, id := range yyDollar[3].strings {
				yyVAL.decls.vars = append(yyVAL.decls.vars, &varId{varDecl: varDecl{name: id, typ: yyDollar[5].ptyp}})
			}
		}
	case 5:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:60
		{
			yyVAL.decls.types = append(yyDollar[1].decls.types, &typTypedef{name: yyDollar[3].s, typ: yyDollar[5].ptyp})
		}
	case 6:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:61
		{
			yyVAL.decls = pDecls{}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:64
		{
			yyVAL.ptyp = yyDollar[1].ptyp
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:65
		{
			yyVAL.ptyp = typTypedef{name: yyDollar[1].s}
		}
	case 9:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:66
		{
			yyVAL.ptyp = typArray{start: yyDollar[3].i, end: yyDollar[5].i, typ: yyDollar[8].ptyp}
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:67
		{
			yyVAL.ptyp = typPointer{typ: yyDollar[2].ptyp}
		}
	case 11:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:68
		{
			yyVAL.ptyp = typRecord{fields: yyDollar[2].vars}
		}
	case 12:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:69
		{
			yyVAL.ptyp = typVoid{}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:70
		{
			yyVAL.ptyp = typVoid{}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:73
		{
			yyVAL.ptyp = typPrimitive{primInt}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:74
		{
			yyVAL.ptyp = typPrimitive{primReal}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:75
		{
			yyVAL.ptyp = typPrimitive{primChar}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:76
		{
			yyVAL.ptyp = typPrimitive{primBool}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:77
		{
			yyVAL.ptyp = typPrimitive{primString}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:80
		{
			yyVAL.functions = append(yyDollar[1].functions, yyDollar[2].function)
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:81
		{
			yyVAL.functions = nil
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:84
		{
			yyDollar[1].function.decls = yyDollar[2].decls.vars
			yyDollar[1].function.body = yyDollar[3].stmt
			yyVAL.function = yyDollar[1].function
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:87
		{
			yyVAL.function = &varFunction{
				varDecl: varDecl{
					name: yyDollar[2].s,
					typ:  typFunction{name: yyDollar[2].s, args: yyDollar[3].vars, ret: yyDollar[5].ptyp},
				},
				args: yyDollar[3].vars,
				ret:  &varId{varDecl: varDecl{typ: yyDollar[5].ptyp}},
			}
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:97
		{
			yyVAL.function = &varFunction{
				varDecl: varDecl{
					name: yyDollar[2].s,
					typ:  typFunction{name: yyDollar[2].s, args: yyDollar[3].vars, ret: typVoid{}},
				},
				args: yyDollar[3].vars,
				ret:  &varId{varDecl: varDecl{typ: typVoid{}}},
			}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:109
		{
			yyVAL.vars = yyDollar[2].vars
		}
	case 25:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:110
		{
			yyVAL.vars = nil
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:113
		{
			yyVAL.vars = append(yyDollar[1].vars, yyDollar[3].vars...)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:114
		{
			yyVAL.vars = yyDollar[1].vars
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:117
		{
			for _, id := range yyDollar[2].strings {
				yyVAL.vars = append(yyVAL.vars, &varId{varDecl: varDecl{name: id, typ: yyDollar[4].ptyp}})
			}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:118
		{
			for _, id := range yyDollar[1].strings {
				yyVAL.vars = append(yyVAL.vars, &varId{varDecl: varDecl{name: id, typ: yyDollar[3].ptyp}})
			}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:121
		{
			yyVAL.stmt = &stmBlock{stmts: yyDollar[2].stmts}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:123
		{
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 32:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:124
		{
			yyVAL.stmts = nil
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:127
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:128
		{
			yyVAL.stmts = []stmt{yyDollar[1].stmt}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:131
		{
			yyVAL.stmt = &stmAssign{id: yyDollar[1].expr, e: yyDollar[3].expr}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:132
		{
			yyVAL.stmt = &stmBreak{}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:133
		{
			yyVAL.stmt = &stmContinue{}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:134
		{
			yyVAL.stmt = &stmCall{fn: &expId{name: yyDollar[1].s}, args: yyDollar[3].exprs}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:135
		{
			yyVAL.stmt = &stmCall{fn: &expId{name: yyDollar[1].s}}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:136
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:137
		{
			yyVAL.stmt = &stmIf{cond: yyDollar[2].expr, ifTrue: yyDollar[4].stmt}
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:138
		{
			yyVAL.stmt = &stmIf{cond: yyDollar[2].expr, ifTrue: yyDollar[4].stmt, ifFalse: yyDollar[6].stmt}
		}
	case 43:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:139
		{
			yyVAL.stmt = &stmFor{counter: &expId{name: yyDollar[2].s}, expr1: yyDollar[4].expr, expr2: yyDollar[6].expr, body: yyDollar[8].stmt}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:140
		{
			yyVAL.stmt = &stmWhile{e: yyDollar[2].expr, body: yyDollar[4].stmt}
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:141
		{
			yyVAL.stmt = &stmRepeat{e: yyDollar[4].expr, body: &stmBlock{yyDollar[2].stmts}}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:144
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:145
		{
			yyVAL.exprs = []expr{yyDollar[1].expr}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:148
		{
			yyVAL.expr = &expCall{fn: &expId{name: yyDollar[1].s}, args: yyDollar[3].exprs}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:149
		{
			yyVAL.expr = &expUnop{op: unopAt, e: yyDollar[2].expr}
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:150
		{
			yyVAL.expr = &expUnop{op: unopPlus, e: yyDollar[2].expr}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:151
		{
			yyVAL.expr = &expUnop{op: unopMinus, e: yyDollar[2].expr}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:152
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:153
		{
			yyVAL.expr = &expBinop{op: binAND, left: yyDollar[1].expr, right: yyDollar[3].expr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:154
		{
			yyVAL.expr = &expBinop{op: binDIV, left: yyDollar[1].expr, right: yyDollar[3].expr}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:155
		{
			yyVAL.expr = &expBinop{op: binLT, left: yyDollar[1].expr, right: yyDollar[3].expr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:156
		{
			yyVAL.expr = &expBinop{op: binEQ, left: yyDollar[1].expr, right: yyDollar[3].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:157
		{
			yyVAL.expr = &expBinop{op: binGT, left: yyDollar[1].expr, right: yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:158
		{
			yyVAL.expr = &expBinop{op: binSUB, left: yyDollar[1].expr, right: yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:159
		{
			yyVAL.expr = &expBinop{op: binFDIV, left: yyDollar[1].expr, right: yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:160
		{
			yyVAL.expr = &expBinop{op: binMUL, left: yyDollar[1].expr, right: yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:161
		{
			yyVAL.expr = &expBinop{op: binADD, left: yyDollar[1].expr, right: yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:162
		{
			yyVAL.expr = &expBinop{op: binGE, left: yyDollar[1].expr, right: yyDollar[3].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:163
		{
			yyVAL.expr = &expBinop{op: binLE, left: yyDollar[1].expr, right: yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:164
		{
			yyVAL.expr = &expBinop{op: binMOD, left: yyDollar[1].expr, right: yyDollar[3].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:165
		{
			yyVAL.expr = &expBinop{op: binNE, left: yyDollar[1].expr, right: yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:166
		{
			yyVAL.expr = &expBinop{op: binOR, left: yyDollar[1].expr, right: yyDollar[3].expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:167
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:168
		{
			yyVAL.expr = &expUnop{op: unopNot, e: yyDollar[2].expr}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:169
		{
			yyVAL.expr = &expConst{t: primInt, i: yyDollar[1].i}
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:170
		{
			yyVAL.expr = &expConst{t: primReal, f: yyDollar[1].f}
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:171
		{
			yyVAL.expr = &expConst{t: primBool, b: true}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:172
		{
			yyVAL.expr = &expConst{t: primBool, b: false}
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:173
		{
			yyVAL.expr = &expConst{t: primString, s: yyDollar[1].s}
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:176
		{
			yyVAL.expr = &expId{name: yyDollar[1].s}
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:177
		{
			yyVAL.expr = &expBinop{op: binArrayIndex, left: yyDollar[1].expr, right: yyDollar[3].expr}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:178
		{
			yyVAL.expr = &expField{e: yyDollar[1].expr, field: &expId{name: yyDollar[3].s}}
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:179
		{
			yyVAL.expr = &expUnop{op: unopPtr, e: yyDollar[1].expr}
		}
	}
	goto yystack /* stack new state and value */
}
