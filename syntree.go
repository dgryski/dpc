package main

type varDecl struct {
	name string
	typ  pType
}

func (p *varDecl) Name() string { return p.name }
func (p *varDecl) Type() pType  { return p.typ }

type varProgram struct {
	varDecl
	vars     []*varId
	types    []*typTypedef
	subprogs []*varFunction
	body     stmt
}

type varId struct {
	varDecl
	byReference bool
}

type varFunction struct {
	varDecl
	args  []*varId
	ret   *varId
	decls []*varId
	body  stmt
}

type varType struct {
	varDecl
}

type pType interface {
	Size() int
	typeNode()
}

type pDecls struct {
	vars  []*varId
	types []*typTypedef
}

type pvariable interface {
	Name() string
	Type() pType
	varNode()
}

func (p *varProgram) varNode()  {}
func (p *varId) varNode()       {}
func (p *varFunction) varNode() {}

type primitive int

//go:generate stringer -type=primitive
const (
	primInt primitive = iota
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
	primType primitive
}

func (p typPrimitive) Size() int { return 8 }

type typRecord struct {
	fields []*varId
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
	args []*varId
	ret  pType
}

func (f typFunction) Size() int { return 8 }

type typTypedef struct {
	name string
	typ  pType
}

func (t typTypedef) Size() int { return t.typ.Size() }

type expr interface {
	IsLValue() bool
	Type() pType
	exprNode()
}

type expConst struct {
	i int
	f float64
	s string
	b bool

	t primitive

	typ pType
}

func (e *expConst) IsLValue() bool { return false }
func (e *expConst) Type() pType    { return e.typ }
func (e *expConst) exprNode()      {}

type expId struct {
	name string

	bound pvariable
	typ   pType
}

func (e *expId) IsLValue() bool { return true }
func (e *expId) Type() pType    { return e.typ }
func (e *expId) exprNode()      {}

type expField struct {
	e      expr
	field  *expId
	record typRecord

	typ pType
}

func (e *expField) IsLValue() bool { return true }
func (e *expField) Type() pType    { return e.typ }
func (e *expField) exprNode()      {}

type expCall struct {
	fn   *expId
	args []expr

	typ pType
}

func (e *expCall) IsLValue() bool { return true }
func (e *expCall) Type() pType    { return e.typ }
func (e *expCall) exprNode()      {}

type unop byte

//go:generate stringer -type=unop
const (
	unopNot unop = iota
	unopPtr
	unopAt
	unopMinus
	unopPlus
	unopIntToReal
)

type expUnop struct {
	op unop
	e  expr

	typ pType
}

func (e *expUnop) IsLValue() bool { return e.op == unopPtr }
func (e *expUnop) Type() pType    { return e.typ }
func (e *expUnop) exprNode()      {}

type binop byte

//go:generate stringer -type=binop
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

	typ pType
}

func (e *expBinop) IsLValue() bool { return e.op == binArrayIndex }
func (e *expBinop) Type() pType    { return e.typ }
func (e *expBinop) exprNode()      {}

type stmt interface {
	stmtNode()
}

type stmAssign struct {
	id expr
	e  expr
}

type stmCall struct {
	fn   *expId
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
	counter *expId
	expr1   expr
	expr2   expr
	body    stmt
}
type stmWhile struct {
	e    expr
	body stmt
}
type stmRepeat struct {
	e    expr
	body stmt
}

type stmBlock struct {
	stmts []stmt
}

func (s *stmAssign) stmtNode()   {}
func (s *stmBreak) stmtNode()    {}
func (s *stmCall) stmtNode()     {}
func (s *stmContinue) stmtNode() {}
func (s *stmBlock) stmtNode()    {}
func (s *stmIf) stmtNode()       {}
func (s *stmFor) stmtNode()      {}
func (s *stmWhile) stmtNode()    {}
func (s *stmRepeat) stmtNode()   {}
