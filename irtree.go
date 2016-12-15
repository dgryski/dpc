package main

import "fmt"

type irnode interface {
	String() string
	kids() []irnode
	build([]irnode) irnode
}

type ireConst struct {
	value int
}

func (ir *ireConst) String() string { return fmt.Sprintf("(const %d)", ir.value) }

func (ir *ireConst) kids() []irnode { return nil }

func (ir *ireConst) build(k []irnode) irnode { return ir }

type ireFConst struct {
	value float64
	name  string
}

func (ir *ireFConst) String() string { return fmt.Sprintf("(fconst %f %q)", ir.value, ir.name) }

func (ir *ireFConst) kids() []irnode { return nil }

func (ir *ireFConst) build(k []irnode) irnode { return ir }

type ireName struct {
	label string
}

func (ir *ireName) String() string { return fmt.Sprintf("(name %q)", ir.label) }

func (ir *ireName) kids() []irnode { return nil }

func (ir *ireName) build(k []irnode) irnode { return ir }

type ireTemp struct {
	temp  string
	where string
}

func (ir *ireTemp) String() string { return fmt.Sprintf("(temp %q)", ir.temp) }

func (ir *ireTemp) kids() []irnode { return nil }

func (ir *ireTemp) build(k []irnode) irnode { return ir }

type ireBinop struct {
	op          binop
	left, right irnode
}

func (ir *ireBinop) String() string { return fmt.Sprintf("(binop %v %v %v)", ir.op, ir.left, ir.right) }

func (ir *ireBinop) kids() []irnode { return []irnode{ir.left, ir.right} }

func (ir *ireBinop) build(k []irnode) irnode { return &ireBinop{op: ir.op, left: k[0], right: k[1]} }

type ireUnop struct {
	op   unop
	expr irnode
}

func (ir *ireUnop) String() string { return fmt.Sprintf("(unop %v %v)", ir.op, ir.expr) }

func (ir *ireUnop) kids() []irnode { return []irnode{ir.expr} }

func (ir *ireUnop) build(k []irnode) irnode { return &ireUnop{op: ir.op, expr: k[0]} }

type ireMem struct {
	expr irnode
}

func (ir *ireMem) String() string { return fmt.Sprintf("(mem %v)", ir.expr) }

func (ir *ireMem) kids() []irnode { return []irnode{ir.expr} }

func (ir *ireMem) build(k []irnode) irnode { return &ireMem{expr: k[0]} }

type ireCall struct {
	function irnode
	args     []irnode
}

func (ir *ireCall) String() string {
	var s string
	for _, a := range ir.args {
		s += fmt.Sprintf(" %s", a)
	}

	return fmt.Sprintf("(call %v (args%s)", ir.function, s)
}

func (ir *ireCall) kids() []irnode {
	var k []irnode
	k = append(k, ir.function)
	return append(k, ir.args...)
}

func (ir *ireCall) build(k []irnode) irnode {
	return &ireCall{
		function: k[0],
		args:     k[1:],
	}
}

type ireESeq struct {
	stmt irnode
	expr irnode
}

func (ir *ireESeq) String() string { return fmt.Sprintf("(eseq %v %v)", ir.stmt, ir.expr) }

func (ir *ireESeq) kids() []irnode { return []irnode{ir.stmt, ir.expr} }

func (ir *ireESeq) build(k []irnode) irnode { return &ireESeq{stmt: k[0], expr: k[1]} }

type irsPrologue struct {
	framesize int
	argcount  int
}

func (ir *irsPrologue) String() string { return "(prologue)" }

func (ir *irsPrologue) kids() []irnode { return nil }

func (ir *irsPrologue) build(k []irnode) irnode { return ir }

type irsEpilogue struct {
	retval irnode
}

func (ir *irsEpilogue) String() string { return "(epilogue)" }

func (ir *irsEpilogue) kids() []irnode { return nil }

func (ir *irsEpilogue) build(k []irnode) irnode { return ir }

type irsMove struct {
	dst, src irnode
}

func (ir *irsMove) String() string { return fmt.Sprintf("(move %v %v)", ir.dst, ir.src) }

func (ir *irsMove) kids() []irnode {
	if mem, ok := ir.dst.(*ireMem); ok {
		return []irnode{mem.expr, ir.src}
	}
	return []irnode{ir.src}
}

func (ir *irsMove) build(k []irnode) irnode {
	if _, ok := ir.dst.(*ireMem); ok {
		return &irsMove{&ireMem{expr: k[0]}, k[1]}
	}
	return &irsMove{ir.dst, k[0]}
}

type irsMoveCall struct {
	dst, call irnode
}

func (ir *irsMoveCall) String() string { return fmt.Sprintf("(movecall %v %v)", ir.dst, ir.call) }

func (ir *irsMoveCall) kids() []irnode {
	return ir.call.kids()
}

func (ir *irsMoveCall) build(k []irnode) irnode {
	return &irsMove{ir.dst, ir.call.build(k)}
}

type irsExpr struct {
	expr irnode
}

func (ir *irsExpr) String() string { return fmt.Sprintf("(expr %v)", ir.expr) }

func (ir *irsExpr) kids() []irnode {
	return []irnode{ir.expr}
}

func (ir *irsExpr) build(k []irnode) irnode {
	return &irsExpr{k[0]}
}

type irsExprCall struct {
	call irnode
}

func (ir *irsExprCall) String() string { return fmt.Sprintf("(exprcall %v)", ir.call) }

func (ir *irsExprCall) kids() []irnode {
	return ir.call.kids()
}

func (ir *irsExprCall) build(k []irnode) irnode {
	return &irsExpr{ir.call.build(k)}
}

type irsJump struct {
	target irnode
}

func (ir *irsJump) String() string { return fmt.Sprintf("(jump %v)", ir.target) }

func (ir *irsJump) kids() []irnode {
	return []irnode{ir.target}
}

func (ir *irsJump) build(k []irnode) irnode {
	return &irsJump{k[0]}
}

type irsCJump struct {
	relop           binop
	left, right     irnode
	iftrue, iffalse irnode
}

func (ir *irsCJump) String() string {
	return fmt.Sprintf("(cjump %v %v %v %v %v)", ir.relop, ir.left, ir.right, ir.iftrue, ir.iffalse)
}

func (ir *irsCJump) kids() []irnode {
	return []irnode{ir.left, ir.right, ir.iftrue, ir.iffalse}
}

func (ir *irsCJump) build(k []irnode) irnode {
	return &irsCJump{ir.relop, k[0], k[1], k[2], k[3]}
}

type irsSeq struct {
	seq []irnode
}

func (ir *irsSeq) String() string {
	var s string
	for _, ss := range ir.seq {
		s += "\n" + ss.String()
	}
	return fmt.Sprintf("(seq %v)", s)
}

func (ir *irsSeq) kids() []irnode {
	return ir.seq
}

func (ir *irsSeq) build(k []irnode) irnode {
	return &irsSeq{k}
}

type irsLabel struct {
	label string
}

func (ir *irsLabel) String() string { return fmt.Sprintf("(label %v)", ir.label) }

func (ir *irsLabel) kids() []irnode {
	return nil
}

func (ir *irsLabel) build(k []irnode) irnode {
	return ir
}

type irsNop struct {
}

func (ir *irsNop) String() string { return "(nop)" }

func (ir *irsNop) kids() []irnode {
	return nil
}

func (ir *irsNop) build(k []irnode) irnode {
	return ir
}
