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
