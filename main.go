package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, _ := ioutil.ReadAll(os.Stdin)
	fmt.Printf("%v\n", lex(data))
}
