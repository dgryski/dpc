//go:generate ragel -Z lexer.rl
//go:generate go tool yacc parser.y

package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	data, _ := ioutil.ReadAll(os.Stdin)
	tokens := lex(data)

	l := pascalLexer(tokens)

	if yyParse(&l) != 0 {
		log.Fatal("parse error")
	}

	log.Printf("%#v\n", program)
}
