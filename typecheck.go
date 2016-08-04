package main

import (
	"fmt"
	"strconv"
)

func typecheck(program varProgram) {

	for _, f := range program.subprogs {
		typecheckStmt(f.body)
	}

	typecheckStmt(program.body)
}

func typecheckStmt(stmt stmt) {

	if stmt == nil {
		return
	}

	switch stmt := stmt.(type) {
	case *stmBreak:
	case *stmContinue:
	case *stmFor:
		typecheckExpr(stmt.counter)
		typecheckExpr(stmt.expr1)
		typecheckExpr(stmt.expr2)
		typecheckStmt(stmt.body)
		if !isPrimitive(stmt.counter.bound.Type(), primInt) {
			fmt.Printf("invalid type %#v for loop variable\n", stmt.counter.bound.Type())
		}
		if !isPrimitive(stmt.expr1.Type(), primInt) {
			fmt.Printf("invalid type %#v for loop start value\n", stmt.expr1.Type())
		}
		if !isPrimitive(stmt.expr2.Type(), primInt) {
			fmt.Printf("invalid type %#v for loop end value\n", stmt.expr2.Type())
		}
	case *stmWhile:
		typecheckExpr(stmt.e)
		if !isPrimitive(stmt.e.Type(), primBool) {
			fmt.Printf("invalid type %#v for while condition\n", stmt.e.Type())
		}
		typecheckStmt(stmt.body)
	case *stmRepeat:
		typecheckExpr(stmt.e)
		typecheckStmt(stmt.body)
		if !isPrimitive(stmt.e.Type(), primBool) {
			fmt.Printf("invalid type %#v for repeat condition\n", stmt.e.Type())
		}
	case *stmIf:
		typecheckExpr(stmt.cond)
		typecheckStmt(stmt.ifTrue)
		typecheckStmt(stmt.ifFalse)
		if !isPrimitive(stmt.cond.Type(), primBool) {
			fmt.Printf("invalid type %#v for if condition\n", stmt.cond.Type())
		}
	case *stmAssign:
		typecheckExpr(stmt.id)
		typecheckExpr(stmt.e)

		if !stmt.id.IsLValue() {
			fmt.Printf("invalid lvalue for assignment")
		}

		// `float := int` is valid
		if isPrimitive(stmt.id.Type(), primReal) && isPrimitive(stmt.e.Type(), primInt) {
			stmt.e = promoteToReal(stmt.e)
		}

		if stmt.id.Type() != stmt.e.Type() {
			fmt.Printf("invalid types for assignment (`%#v` and `%#v`)", stmt.id.Type(), stmt.e.Type())
		}

	case *stmBlock:
		for _, ss := range stmt.stmts {
			typecheckStmt(ss)
		}
	case *stmCall:
		typecheckExpr(stmt.fn)
		for _, e := range stmt.args {
			typecheckExpr(e)
		}

		// typecheckBuiltin(...)
		typecheckCallStmt(stmt)
	default:
		panic("unknown stmt: " + fmt.Sprintf("%T", stmt))
	}
}

func typecheckExpr(e expr) {
	switch e := e.(type) {
	case *expConst:
		e.typ = typPrimitive{primType: e.t}
	case *expId:
		if e.bound == nil {
			fmt.Printf("unknown identifier %#v", e.name)
		} else {
			e.typ = e.bound.Type()
		}
	case *expBinop:
		typecheckBinop(e)
	case *expCall:
		typecheckCallExpr(e)
	case *expUnop:
		typecheckUnop(e)
	case *expField:
		// TODO(dgryski): what goes here?
		// typecheckExpr(e)
		e.typ = e.field.bound.Type()
	default:
		panic("unknown expr: " + fmt.Sprintf("%T", e))
	}
}

func typecheckCallExpr(expr expr) {
	// TODO(dgryski): typecheckCallExpr
}

func typecheckCallStmt(stmt stmt) {
	// TODO(dgryski): typecheckCallStmt
}

func typecheckBinop(e *expBinop) {
	typecheckExpr(e.left)
	typecheckExpr(e.right)

	switch e.op {
	case binAND, binOR:
		e.typ = e.left.Type()
		if !isPrimitive(e.left.Type(), primBool) || !isPrimitive(e.right.Type(), primBool) {
			fmt.Println("boolean types required")
		}

	case binDIV, binMOD:
		e.typ = e.left.Type()
		if !isPrimitive(e.left.Type(), primInt) || !isPrimitive(e.right.Type(), primInt) {
			fmt.Println("integer types required")
		}

	case binFDIV:
		e.typ = typPrimitive{primType: primReal}
		if !isArithmetic(e.left.Type()) || !isArithmetic(e.right.Type()) {
			fmt.Println("arithmetic types required")
		}

		e.left, e.right = arithmeticPromotion(e.left, e.right)

	case binADD, binSUB, binMUL:
		e.typ = e.left.Type()
		if !isArithmetic(e.left.Type()) || !isArithmetic(e.right.Type()) {
			fmt.Println("arithmetic types required")
		}

		e.left, e.right = arithmeticPromotion(e.left, e.right)

	case binLT, binLE, binEQ, binGE, binGT, binNE:

		if isArithmetic(e.left.Type()) && isArithmetic(e.right.Type()) {
			e.left, e.right = arithmeticPromotion(e.left, e.right)
		} else if isPrimitive(e.left.Type(), primBool) && isPrimitive(e.right.Type(), primBool) {
			// nothing
		} else {
			fmt.Printf("invalid types `%#v` and `%#v` to relational operator", e.left.Type(), e.right.Type())
		}

		e.typ = typPrimitive{primType: primBool}

	case binArrayIndex:
		fmt.Printf("e %+v\n", e)
		arr, ok := e.left.Type().(typArray)
		if !ok {
			fmt.Printf("invalid type `%#v' for array in array index", e.left)
		}

		if !isPrimitive(e.right.Type(), primInt) {
			fmt.Printf("invalid type `%#v' for index in array index", e.right)
		}

		e.typ = arr.typ

	default:
		panic("unknown binop: " + strconv.Itoa(int(e.op)))
	}
}

func typecheckUnop(e *expUnop) {
	typecheckExpr(e.e)

	switch e.op {
	case unopPlus, unopMinus:
		e.typ = e.e.Type()
	case unopPtr:
		ptr, ok := e.e.Type().(typPointer)
		if !ok {
			fmt.Printf("invalid type `%#v' for pointer dereference", e.e.Type())
		}

		e.typ = ptr.typ
	case unopAt:
		if !e.e.IsLValue() {
			fmt.Printf("bad lvalue for address-of")
		}
		e.typ = typPointer{typ: e.typ}
	case unopNot:
		if !isPrimitive(e.e.Type(), primBool) {
			fmt.Printf("boolean type required")
		}
		e.typ = e.e.Type()
	default:
		panic("unknown unary operator")
	}
}

func isPrimitive(t pType, want primitive) bool {
	prim, ok := t.(typPrimitive)
	if !ok {
		return false
	}

	return prim.primType == want
}

func isArithmetic(t pType) bool {
	prim, ok := t.(typPrimitive)
	if !ok {
		return false
	}

	return prim.primType == primInt || prim.primType == primReal
}

func promoteToReal(e expr) expr {

	// ok to panic here
	p := e.Type().(typPrimitive)

	switch p.primType {
	case primReal:
		return e
	case primInt:
		return &expUnop{op: unopIntToReal, typ: typPrimitive{primReal}, e: e}
	default:
		panic("bad type in promoteToReal: " + strconv.Itoa(int(p.primType)))
	}
}

func arithmeticPromotion(left, right expr) (expr, expr) {

	// ok to panic()
	lt := left.Type().(typPrimitive).primType
	rt := right.Type().(typPrimitive).primType

	switch {
	case lt == primInt && rt == primInt:
		return left, right
	case lt == primReal && rt == primReal:
		return left, right
	case lt == primInt && rt == primReal:
		return promoteToReal(left), right
	case lt == primReal && rt == primInt:
		return left, promoteToReal(right)
	}

	panic("unreached")
}

func callReturnType(b pvariable) pType {

	switch v := b.(type) {
	case *varFunction:
		return v.ret.typ
	case *varId:
		vptr, ok := v.typ.(*typPointer)
		if !ok {
			return nil
		}

		vfn, ok := vptr.typ.(*typFunction)
		if !ok {
			return nil
		}

		return vfn.ret
	}

	panic("unknown return type")
}
