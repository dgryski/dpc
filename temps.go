package main

import "fmt"

var labelCount int

func newTempLabel(l string) irnode {
	if l == "" {
		l = fmt.Sprintf("L%d", labelCount)
	}

	return &irsLabel{l}
}

var regCount int

func newTempReg(r string) irnode {
	if r == "" {
		r = fmt.Sprintf("t%d", regCount)
	}

	return &ireTemp{r, ""}
}

var nameCount int

func newTempName(n string) irnode {
	if n == "" {
		n = fmt.Sprintf("LC%d", nameCount)
	}

	return &ireName{n}
}
