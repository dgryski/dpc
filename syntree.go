package main

type pProgram struct {
	name     string
	vars     []pVar
	types    []typTypedef
	subprogs []pFunction
}

type pVar struct {
	name        string
	typ         pType
	byReference bool
}

type pFunction struct {
	name  string
	args  []pVar
	ret   pType
	decls []pVar
}

type pType interface {
	Size() int
	typeNode()
}

type pDecls struct {
	vars  []pVar
	types []typTypedef
}

type Primitive int

const (
	primInt Primitive = iota
	primBool
	primReal
	primString
	primChar
)

type typVoid struct {
}

func (v typVoid) Size() int { return 0 }

func (_ typArray) typeNode()     {}
func (_ typFunction) typeNode()  {}
func (_ typPrimitive) typeNode() {}
func (_ typPointer) typeNode()   {}
func (_ typRecord) typeNode()    {}
func (_ typVoid) typeNode()      {}
func (_ typTypedef) typeNode()   {}

type typPrimitive struct {
	primtType Primitive
}

func (p typPrimitive) Size() int { return 8 }

type typRecord struct {
	fields []pVar
}

func (r typRecord) Size() int {
	var size int
	for _, f := range r.fields {
		size += f.typ.Size()
	}
	return size
}

type typPointer struct {
	typ pType
}

func (p typPointer) Size() int { return 8 }

type typArray struct {
	start, end int
	typ        pType
}

func (a typArray) Size() int { return (a.end - a.start) * a.typ.Size() }

type typFunction struct {
	name string
	args []pVar
	ret  pType
}

type typTypedef struct {
	name string
	typ  pType
}

func (t typTypedef) Size() int { return t.typ.Size() }

type expr interface {
	IsLValue()
	exprNode()
}

type expConst struct {
	i int
	f float64
	s string

	t pType
}

func (c expConst) IsLValue() bool { return false }

type expId struct {
	name  string
	byRef bool

	// bound

}

func (c expId) IsLValue() bool { return true }

type expField struct {
	record typRecord
	field  pVar
}

func (c expField) IsLValue() bool { return true }

type expCall struct {
	fn   expr
	args []expr
}

func (c expCall) IsLValue() bool { return true }

type unop byte

const (
	unopNot unop = iota
	unopPtr
	unopAt
)

type expUnop struct {
	op unop
	e  expr
}

type binop byte

const (
	binAND binop = iota
	binDIV
	binLT
	binEQ
	binGT
	binSUB
	binFDIV
	binMUL
	binADD
	binGE
	binLE
	binMOD
	binNE
	binOR
)

type expBinop struct {
	op          binop
	left, right expr
}
