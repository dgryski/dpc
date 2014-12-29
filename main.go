//go:generate ragel -Z lexer.rl
//go:generate go tool yacc parser.y

package main

import (
	"io/ioutil"
	"os"
)

func main() {
	data, _ := ioutil.ReadAll(os.Stdin)
	tokens := lex(data)

	l := pascalLexer(tokens)

	yyParse(&l)
}
