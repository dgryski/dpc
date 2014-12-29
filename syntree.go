package main

type varProgram struct {
	name     string
	vars     []varId
	types    []typTypedef
	subprogs []varFunction
	body     []stmt
}

type varId struct {
	name        string
	typ         pType
	byReference bool
}

type varFunction struct {
	name  string
	args  []varId
	ret   varId
	decls []varId
	body  []stmt
}

type pType interface {
	Size() int
	typeNode()
}

type pDecls struct {
	vars  []varId
	types []typTypedef
}

type pvariable interface {
	varNode()
}

func (p varProgram) varNode()  {}
func (p varId) varNode()       {}
func (p varFunction) varNode() {}

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
	fields []varId
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
	args []varId
	ret  pType
}

type typTypedef struct {
	name string
	typ  pType
}

func (t typTypedef) Size() int { return t.typ.Size() }

type expr interface {
	IsLValue() bool
	exprNode()
}

type expConst struct {
	i int
	f float64
	s string
	b bool

	t Primitive
}

func (e expConst) IsLValue() bool { return false }
func (e expConst) exprNode()      {}

type exvarId struct {
	name  string
	byRef bool

	// bound

}

func (e exvarId) IsLValue() bool { return true }
func (e exvarId) exprNode()      {}

type expField struct {
	e      expr
	record typRecord
	field  varId
}

func (e expField) IsLValue() bool { return true }
func (e expField) exprNode()      {}

type expCall struct {
	fn   varId
	args []expr
}

func (e expCall) IsLValue() bool { return true }
func (e expCall) exprNode()      {}

type unop byte

const (
	unopNot unop = iota
	unopPtr
	unopAt
	unopMinus
	unopPlus
)

type expUnop struct {
	op unop
	e  expr
}

func (e expUnop) IsLValue() bool { return e.op == unopPtr }
func (e expUnop) exprNode()      {}

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
	binArrayIndex
)

type expBinop struct {
	op          binop
	left, right expr
}

func (e expBinop) IsLValue() bool { return false }
func (e expBinop) exprNode()      {}

type stmt interface {
	stmtNode()
}

type stmAssign struct {
	id expr
	e  expr
}

type stmCall struct {
	fn   varId
	args []expr
}

type stmBreak struct{}
type stmContinue struct{}
type stmIf struct {
	cond    expr
	ifTrue  stmt
	ifFalse stmt
}
type stmFor struct {
}
type stmWhile struct {
	e    expr
	body stmt
}
type stmRepeat struct {
	e    expr
	body []stmt
}

type stmBlock struct {
	stmts []stmt
}

func (s stmAssign) stmtNode()   {}
func (s stmBreak) stmtNode()    {}
func (s stmCall) stmtNode()     {}
func (s stmContinue) stmtNode() {}
func (s stmBlock) stmtNode()    {}
func (s stmIf) stmtNode()       {}
func (s stmFor) stmtNode()      {}
func (s stmWhile) stmtNode()    {}
func (s stmRepeat) stmtNode()   {}
