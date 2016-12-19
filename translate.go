package main

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

	panic("unhandled stmt type")
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
	return &irsNop{}
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
	return &irsNop{}
}

func translateExprBinop(e *expBinop) irnode {
	return &irsNop{}
}
func translateExprUnop(e *expUnop) irnode {
	return &irsNop{}
}
func translateExprId(e *expId) irnode {
	return &irsNop{}
}
func translateExprCall(e *expCall) irnode {
	return &irsNop{}
}
func translateExprField(e *expField) irnode {
	return &irsNop{}
}
