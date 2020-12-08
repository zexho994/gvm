package runtime

// 映射到java中的一个thread todo
type Thread struct {
	pc    *ProgramCounter
	stack *Stack
}

func (t *Thread) PC() int {
	return t.pc.pc()
}

func (t *Thread) SetPC(newPc int) {
	t.pc.setpc(newPc)
}

func (t *Thread) PushFrame(newFrame *Frame) {
	t.stack.Push(newFrame)
}

func (t *Thread) PopFrame() *Frame {
	return t.stack.Pop()
}

func (t *Thread) Frame() *Frame {
	return t.stack.Peek()
}

func newThread() *Thread {
	return &Thread{
		stack: NewStack(1024),
	}
}
