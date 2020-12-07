package runtime

type ProgramCounter struct {
	// 当前位置
	cur int
}

func (pc *ProgramCounter) pc() int {
	return pc.cur
}

func (pc *ProgramCounter) setpc(newPc int) {
	pc.cur = newPc
}
