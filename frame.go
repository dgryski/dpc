package main

const pointerSize = 4

type global struct {
	label irnode
	size  int
}

type frame struct {
	framePointer irnode
	stackPointer irnode
	returnValue  irnode
	ret          irnode
	globs        []global
	stringTable  map[string]irnode

	paramOffset  int
	localOffset  int
	return_value ireTemp
}

func (f *frame) reset() {
	f.paramOffset = 4
	f.localOffset = 0
}

func (f *frame) FP() irnode {
	if f.framePointer == nil {
		f.framePointer = newTempReg("fp")
	}

	return f.framePointer
}

func (f *frame) SP() irnode {
	if f.stackPointer == nil {
		f.stackPointer = newTempReg("sp")
	}

	return f.stackPointer
}

func (f *frame) Ret() irnode {
	if f.returnValue == nil {
		f.returnValue = newTempReg("ret")
	}

	return f.returnValue
}

func (f *frame) allocateGlobal(name string, size int) irnode {
	n := &irsLabel{name}
	f.globs = append(f.globs, global{n, size})
	return &ireMem{n}

}

func (f *frame) allocateLocal(size int) irnode {
	f.localOffset += size
	return &ireMem{&ireBinop{binSUB, f.FP(), &ireConst{f.localOffset}}}
}

func (f *frame) allocateParam(size int) irnode {
	f.localOffset += size
	return &ireMem{&ireBinop{binADD, f.FP(), &ireConst{f.paramOffset}}}
}

func (f *frame) allocateString(s string) irnode {
	where, ok := f.stringTable[s]
	if !ok {
		t := newTempName("")
		f.stringTable[s] = t
		f.globs = append(f.globs, global{t, 0})
		where = t
	}

	return where
}

func (f *frame) allocateVars(program varProgram) {
	for _, v := range program.vars {
		v.where = f.allocateGlobal(v.Name(), v.Type().Size())
	}

	for _, subprog := range program.subprogs {
		f.reset()

		for _, v := range subprog.args {
			v.where = f.allocateParam(v.Type().Size())
		}

		for _, v := range subprog.decls {
			v.where = f.allocateLocal(v.Type().Size())
		}
	}
}
