package main

import "fmt"

type exprLabels struct {
	iftrue  *irsLabel
	iffalse *irsLabel
}

type stmtLabels struct {
	lbreak    *irsLabel
	lcontinue *irsLabel
}

func translateProgram(program varProgram) irnode {

	var seq []irnode

	for _, subprog := range program.subprogs {
		seq = append(seq,
			&irsLabel{subprog.name},
			&irsPrologue{subprog.frameSize, len(subprog.args)},
			translateStmt(subprog.body, nil),
			&irsEpilogue{subprog.ret.where},
		)
	}

	seq = append(seq, &irsLabel{"main"},
		&irsPrologue{4, 3},
		translateStmt(program.body, nil),
		&irsEpilogue{nil},
	)
	return &irsSeq{seq}
}

func translateStmt(body stmt, l *stmtLabels) irnode {

	if body == nil {
		return &irsNop{}
	}

	switch s := body.(type) {
	case *stmAssign:
		return translateStmtAssign(s, l)
	case *stmBreak:
		return translateStmtBreak(s, l)
	case *stmCall:
		return translateStmtCall(s, l)
	case *stmContinue:
		return translateStmtContinue(s, l)
	case *stmBlock:
		return translateStmtBlock(s, l)
	case *stmIf:
		return translateStmtIf(s, l)
	case *stmFor:
		return translateStmtFor(s, l)
	case *stmWhile:
		return translateStmtWhile(s, l)
	case *stmRepeat:
		return translateStmtRepeat(s, l)
	}

	panic(fmt.Sprintf("unhandled stmt type: %v", body))
}

func translateStmtAssign(s *stmAssign, l *stmtLabels) irnode {
	lval := translateExpr(s.id)
	rval := translateExpr(s.e)
	return &irsMove{lval, rval}
}

func translateStmtBreak(s *stmBreak, l *stmtLabels) irnode {
	if l == nil {
		panic("break outside loop")
	}
	return &irsJump{l.lbreak}
}
func translateStmtCall(s *stmCall, l *stmtLabels) irnode {
	return &irsNop{}
}
func translateStmtContinue(s *stmContinue, l *stmtLabels) irnode {
	if l == nil {
		panic("continue outside loop")
	}
	return &irsJump{l.lcontinue}
}
func translateStmtBlock(s *stmBlock, l *stmtLabels) irnode {
	var seq []irnode

	for _, ss := range s.stmts {
		seq = append(seq, translateStmt(ss, l))
	}
	return &irsSeq{seq}

}
func translateStmtIf(s *stmIf, l *stmtLabels) irnode {

	labelTrue := newTempLabel("")
	labelFalse := newTempLabel("")
	labelAfter := newTempLabel("")

	var exprLabels = &exprLabels{iftrue: labelTrue, iffalse: labelFalse}

	return &irsSeq{[]irnode{
		translateCExpr(s.cond, exprLabels),
		labelTrue,
		translateStmt(s.ifTrue, l),
		&irsJump{labelAfter},
		translateStmt(s.ifFalse, l),
		&irsJump{labelAfter},
		labelAfter,
	}}
}

func translateStmtFor(s *stmFor, l *stmtLabels) irnode {

	labelTest := newTempLabel("")
	labelBody := newTempLabel("")
	labelAfter := newTempLabel("")
	labelIncr := newTempLabel("")

	counter := translateExpr(s.counter)
	expr2 := translateExpr(s.expr2)

	slabels := &stmtLabels{lbreak: labelAfter, lcontinue: labelIncr}

	var code []irnode
	code = append(code,
		&irsMove{counter, translateExpr(s.expr1)},
	)

	if _, ok := expr2.(*ireConst); ok {
		code = append(code,
			labelTest,
			&irsCJump{binLE, counter, expr2, labelBody, labelAfter},
		)
	} else {
		regFini := newTempReg("")
		code = append(code,
			&irsMove{regFini, expr2},
			labelTest,
			&irsCJump{binLE, counter, regFini, labelBody, labelAfter},
		)
	}

	code = append(code,
		labelBody,
		translateStmt(s.body, slabels),
		labelIncr,
		&irsMove{counter, &ireBinop{binADD, &ireConst{0}, counter}},
		&irsJump{labelTest},
		labelAfter,
	)
	return &irsSeq{code}
}
func translateStmtWhile(s *stmWhile, l *stmtLabels) irnode {
	labelLoop := newTempLabel("")
	labelAfter := newTempLabel("")
	labelTest := newTempLabel("")

	var exprLabels = &exprLabels{iftrue: labelLoop, iffalse: labelAfter}
	var stmtLabels = &stmtLabels{lbreak: labelAfter, lcontinue: labelTest}

	return &irsSeq{[]irnode{
		&irsJump{labelTest},
		labelLoop,
		translateStmt(s.body, stmtLabels),
		labelTest,
		&irsExpr{translateCExpr(s.e, exprLabels)},
		labelAfter,
	}}
}
func translateStmtRepeat(s *stmRepeat, l *stmtLabels) irnode {
	labelLoop := newTempLabel("")
	labelAfter := newTempLabel("")
	labelTest := newTempLabel("")

	var exprLabels = &exprLabels{iffalse: labelLoop, iftrue: labelAfter}
	var stmtLabels = &stmtLabels{lbreak: labelAfter, lcontinue: labelTest}

	return &irsSeq{[]irnode{
		labelLoop,
		translateStmt(s.body, stmtLabels),
		labelTest,
		&irsExpr{translateCExpr(s.e, exprLabels)},
		labelAfter,
	}}
}

func translateCExpr(e expr, l *exprLabels) irnode {

	if l == nil {
		panic("missing labels")
	}

	switch e.(type) {
	case *expBinop:
		return translateCExprBinop(e, l)
	case *expUnop:
		return translateCExprUnop(e, l)
	}

	exprcode := translateExpr(e)

	return &ireESeq{
		&irsCJump{binNE, &ireConst{0}, exprcode, l.iftrue, l.iffalse},
		&ireConst{0},
	}
}

func translateCExprBinop(e expr, l *exprLabels) irnode {
	return &irsNop{}
}

func translateCExprUnop(e expr, l *exprLabels) irnode {
	return &irsNop{}
}

func translateExpr(e expr) irnode {

	switch e := e.(type) {
	case *expConst:
		return translateExprConst(e)
	case *expBinop:
		return translateExprBinop(e)
	case *expUnop:
		return translateExprUnop(e)
	case *expId:
		return translateExprId(e)
	case *expCall:
		return translateExprCall(e)
	case *expField:
		return translateExprField(e)
	}
	panic("unhandled expr type")
}

func translateExprConst(e *expConst) irnode {
	switch e.t {
	case primBool:
		if e.b {
			return &ireConst{1}
		}
		return &ireConst{0}

	case primInt:
		return &ireConst{e.i}
	}
	panic("unhandled const")
}

func translateExprBinop(e *expBinop) irnode {

	switch e.op {
	case binLE, binLT, binGT, binGE, binNE, binEQ:

		result := newTempReg("")
		iftrue := newTempLabel("")
		iffalse := newTempLabel("")

		return &ireESeq{
			&irsSeq{[]irnode{
				&irsMove{result, &ireConst{0}},
				&irsCJump{e.op, translateExpr(e.left), translateExpr(e.right), iftrue, iffalse},
				iftrue,
				&irsMove{result, &ireConst{1}},
				iffalse,
			}},
			result,
		}
	case binADD, binFDIV, binMUL, binSUB, binMOD, binDIV:
		return &ireBinop{e.op, translateExpr(e.left), translateExpr(e.right)}

	case binOR:
		btest := newTempLabel("")
		result := newTempReg("")
		setResult := newTempLabel("")
		out := newTempLabel("")

		return &ireESeq{
			&irsSeq{[]irnode{
				&irsMove{result, &ireConst{0}},
				&irsCJump{binNE, &ireConst{0}, translateExpr(e.left), setResult, btest},
				btest,
				&irsCJump{binNE, &ireConst{0}, translateExpr(e.right), setResult, out},
				setResult,
				&irsMove{result, &ireConst{1}},
				out,
			}},
			result,
		}
	case binAND:
		btest := newTempLabel("")
		out := newTempLabel("")

		result := newTempReg("")
		setResult := newTempLabel("")

		return &ireESeq{
			&irsSeq{[]irnode{
				&irsMove{result, &ireConst{0}},
				&irsCJump{binEQ, &ireConst{0}, translateExpr(e.left), out, btest},
				btest,
				&irsCJump{binEQ, &ireConst{0}, translateExpr(e.right), out, setResult},
				setResult,
				&irsMove{result, &ireConst{1}},
				out,
			}},
			result,
		}
	case binArrayIndex:
		arr := translateExpr(e.left)

		if _, ok := arr.(*ireMem); !ok {
			panic("array not memory reference")
		}

		arrTyp := e.left.Type().(typArray)

		idx := translateExpr(e.right)

		// a[i] => *(a + (i-start)*sizeof(a[0]))
		//      => *(a + i*sizeof(a[0]) - start*sizeof(a[0]))
		return &ireMem{
			&ireBinop{binADD,
				arr,
				&ireBinop{binSUB,
					&ireBinop{binMUL, idx, &ireConst{arrTyp.typ.Size()}},
					&ireBinop{binMUL, &ireConst{arrTyp.start}, &ireConst{arrTyp.typ.Size()}},
				},
			},
		}
	}

	panic("unhandled binop: " + e.op.String())

}

func translateExprUnop(e *expUnop) irnode {
	return &irsNop{}
}
func translateExprId(e *expId) irnode {

	varid := e.bound.(*varId)
	if varid.byReference {
		return &ireMem{varid.where}
	}

	return varid.where
}
func translateExprCall(e *expCall) irnode {
	return &irsNop{}
}
func translateExprField(e *expField) irnode {
	record := translateExpr(e.e)
	offs := e.field.bound.(*varId).where
	return &ireMem{&ireBinop{binADD, record, offs}}
}
