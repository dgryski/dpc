
//line lexer.rl:1
package main

import (
	"fmt"
	"bytes"
	"strconv"
)


//line lexer.rl:10

//line lexer.go:15
var _scanner_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 1, 9, 1, 10, 1, 11, 
	1, 12, 1, 13, 1, 14, 1, 15, 
	1, 16, 1, 17, 1, 18, 
}

var _scanner_key_offsets []byte = []byte{
	0, 0, 1, 3, 4, 29, 30, 33, 
	35, 36, 38, 39, 46, 
}

var _scanner_trans_keys []byte = []byte{
	39, 48, 57, 125, 9, 10, 32, 39, 
	46, 58, 60, 62, 64, 91, 95, 123, 
	125, 40, 47, 48, 57, 59, 61, 65, 
	90, 93, 94, 97, 122, 46, 46, 48, 
	57, 48, 57, 61, 61, 62, 61, 95, 
	48, 57, 65, 90, 97, 122, 125, 
}

var _scanner_single_lengths []byte = []byte{
	0, 1, 0, 1, 13, 1, 1, 0, 
	1, 2, 1, 1, 1, 
}

var _scanner_range_lengths []byte = []byte{
	0, 0, 1, 0, 6, 0, 1, 1, 
	0, 0, 0, 3, 0, 
}

var _scanner_index_offsets []byte = []byte{
	0, 0, 2, 4, 6, 26, 28, 31, 
	33, 35, 38, 40, 45, 
}

var _scanner_trans_targs []byte = []byte{
	4, 1, 7, 4, 4, 3, 4, 4, 
	4, 1, 5, 8, 9, 10, 4, 4, 
	11, 12, 4, 4, 6, 4, 11, 4, 
	11, 0, 4, 4, 2, 6, 4, 7, 
	4, 4, 4, 4, 4, 4, 4, 4, 
	11, 11, 11, 11, 4, 4, 3, 4, 
	4, 4, 4, 4, 4, 4, 4, 4, 
	4, 
}

var _scanner_trans_actions []byte = []byte{
	19, 0, 0, 37, 21, 0, 23, 25, 
	23, 0, 0, 0, 0, 0, 7, 7, 
	0, 5, 7, 7, 5, 7, 0, 7, 
	0, 0, 11, 27, 0, 5, 31, 0, 
	29, 9, 27, 15, 17, 27, 13, 27, 
	0, 0, 0, 0, 33, 21, 0, 37, 
	35, 27, 31, 29, 27, 27, 27, 33, 
	27, 
}

var _scanner_to_state_actions []byte = []byte{
	0, 0, 0, 0, 1, 0, 0, 0, 
	0, 0, 0, 0, 0, 
}

var _scanner_from_state_actions []byte = []byte{
	0, 0, 0, 0, 3, 0, 0, 0, 
	0, 0, 0, 0, 0, 
}

var _scanner_eof_trans []byte = []byte{
	0, 0, 48, 49, 0, 57, 51, 52, 
	57, 57, 57, 56, 57, 
}

const scanner_start int = 4
const scanner_first_final int = 4
const scanner_error int = 0

const scanner_en_main int = 4


//line lexer.rl:11

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

	
//line lexer.go:132
	{
	cs = scanner_start
	ts = 0
	te = 0
	act = 0
	}

//line lexer.go:140
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_acts = int(_scanner_from_state_actions[cs])
	_nacts = uint(_scanner_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _scanner_actions[_acts - 1] {
		case 1:
//line NONE:1
ts = p

//line lexer.go:163
		}
	}

	_keys = int(_scanner_key_offsets[cs])
	_trans = int(_scanner_index_offsets[cs])

	_klen = int(_scanner_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _scanner_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _scanner_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_scanner_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _scanner_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _scanner_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
_eof_trans:
	cs = int(_scanner_trans_targs[_trans])

	if _scanner_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_scanner_trans_actions[_trans])
	_nacts = uint(_scanner_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _scanner_actions[_acts-1] {
		case 2:
//line NONE:1
te = p+1

		case 3:
//line lexer.rl:47
te = p+1
{ add(int(data[ts])) }
		case 4:
//line lexer.rl:48
te = p+1
{ add(tASSIGN) }
		case 5:
//line lexer.rl:49
te = p+1
{ add(tDOTDOT) }
		case 6:
//line lexer.rl:50
te = p+1
{ add(tGE) }
		case 7:
//line lexer.rl:51
te = p+1
{ add(tLE) }
		case 8:
//line lexer.rl:52
te = p+1
{ add(tNE) }
		case 9:
//line lexer.rl:53
te = p+1
{ addstr(tQSTRING, string(data[ts:te])) }
		case 10:
//line lexer.rl:54
te = p+1
{ lineno += bytes.Count(data[ts:te], []byte{'\n'}) }
		case 11:
//line lexer.rl:55
te = p+1
{ }
		case 12:
//line lexer.rl:56
te = p+1
{ lineno++ }
		case 13:
//line lexer.rl:47
te = p
p--
{ add(int(data[ts])) }
		case 14:
//line lexer.rl:57
te = p
p--
{
		    f, _ := strconv.ParseFloat(string(data[ts:te]), 64)
		    addf(tFNUMBER, f)
		}
		case 15:
//line lexer.rl:61
te = p
p--
{
		    i, _ := strconv.Atoi(string(data[ts:te]))
		    addint(tNUMBER, i)
		}
		case 16:
//line lexer.rl:65
te = p
p--
{
			t, ok := reserved[string(data[ts:te])];
			if !ok {
				t = tID
			};
			addstr(t, string(data[ts:te]))
		}
		case 17:
//line lexer.rl:47
p = (te) - 1
{ add(int(data[ts])) }
		case 18:
//line lexer.rl:61
p = (te) - 1
{
		    i, _ := strconv.Atoi(string(data[ts:te]))
		    addint(tNUMBER, i)
		}
//line lexer.go:319
		}
	}

_again:
	_acts = int(_scanner_to_state_actions[cs])
	_nacts = uint(_scanner_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _scanner_actions[_acts-1] {
		case 0:
//line NONE:1
ts = 0

//line lexer.go:333
		}
	}

	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		if _scanner_eof_trans[cs] > 0 {
			_trans = int(_scanner_eof_trans[cs] - 1)
			goto _eof_trans
		}
	}

	_out: {}
	}

//line lexer.rl:76


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

