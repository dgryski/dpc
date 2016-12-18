//go:generate ragel -Z lexer.rl
//go:generate goyacc parser.y

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	data, _ := ioutil.ReadAll(os.Stdin)
	tokens := lex(data)

	l := pascalLexer(tokens)

	if yyParse(&l) != 0 {
		fmt.Println("parse error")
		return
	}

	var f frame

	f.allocateVars(program)

	bind(program)
	typecheck(program)
	fmt.Println(spew.Sdump(program))
}
