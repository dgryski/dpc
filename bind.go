package main

import (
	"fmt"
)

func bind(program varProgram) {

	var scope scopes
	scope.begin()

	// TODO(dgryski): define builtins

	for _, v := range program.vars {
		bindDecl(v, scope)
	}

	for _, v := range program.types {
		bindType(v, scope)
	}

	for _, f := range program.subprogs {
		bindFunction(f, scope)
	}

	bindStmt(program.body, scope)

	scope.end()
}

func bindDecl(v pvariable, scope scopes) {
	scope.define(v.Name(), v)
	bindType(v.Type(), scope)
}

func bindType(typ pType, scope scopes) {

	switch t := typ.(type) {
	case typPrimitive:
	case typVoid:
	case typPointer:
		bindType(t.typ, scope)
	case typArray:
		bindType(t.typ, scope)
	case typRecord:
		scope.begin()
		for _, v := range t.fields {
			bindDecl(v, scope)
		}
		scope.end()
	default:
		panic("unknown type: " + fmt.Sprintf("%T", t))
	}
}

func bindFunction(f *varFunction, scope scopes) {
	scope.define(f.name, f)
	scope.begin()
	scope.define(f.name+"$ret", f.ret)
	for _, v := range f.args {
		bindDecl(v, scope)
	}
	for _, v := range f.decls {
		bindDecl(v, scope)
	}
	bindStmt(f.body, scope)
	scope.end()
}

func bindStmt(stmt stmt, scope scopes) {

	if stmt == nil {
		return
	}

	switch stmt := stmt.(type) {
	case *stmBreak:
	case *stmContinue:
	case *stmFor:
		bindExpr(stmt.counter, scope)
		bindExpr(stmt.expr1, scope)
		bindExpr(stmt.expr2, scope)
		bindStmt(stmt.body, scope)
	case *stmWhile:
		bindExpr(stmt.e, scope)
		bindStmt(stmt.body, scope)
	case *stmRepeat:
		bindExpr(stmt.e, scope)
		bindStmt(stmt.body, scope)
	case *stmIf:
		bindExpr(stmt.cond, scope)
		bindStmt(stmt.ifTrue, scope)
		bindStmt(stmt.ifFalse, scope)
	case *stmAssign:
		bindExpr(stmt.id, scope)
		bindExpr(stmt.e, scope)
	case *stmBlock:
		for _, ss := range stmt.stmts {
			bindStmt(ss, scope)
		}
	case *stmCall:
		bindExpr(stmt.fn, scope)
		for _, e := range stmt.args {
			bindExpr(e, scope)
		}
	default:
		panic("unknown stmt: " + fmt.Sprintf("%T", stmt))
	}
}

func bindExpr(e expr, scope scopes) {
	switch e := e.(type) {
	case *expConst:
	case *expId:
		e.bound = scope.lookup(e.name)
	case *expBinop:
		bindExpr(e.left, scope)
		bindExpr(e.right, scope)
	case *expCall:
		bindExpr(e.fn, scope)
		for _, arg := range e.args {
			bindExpr(arg, scope)
		}
	case *expUnop:
		bindExpr(e.e, scope)
	case *expField:
		// TODO(dgryski): what goes here?
		// bindExpr(e.record, scope)
	default:
		panic("unknown expr: " + fmt.Sprintf("%T", e))
	}
}

type scope map[string]pvariable

type scopes []scope

// begin a scope
func (s *scopes) begin() {
	*s = append(*s, make(scope))
}

// end a scope
func (s *scopes) end() {
	*s = (*s)[:len(*s)-1]
}

// lookup a symbol in the current scope
func (s *scopes) lookup(name string) pvariable {

	for i := len(*s) - 1; i >= 0; i-- {
		v, ok := (*s)[i][name]
		if ok {
			return v
		}
	}

	return nil
}

// define symbol in current scope
func (s *scopes) define(name string, v pvariable) {
	(*s)[len(*s)-1][name] = v
}
