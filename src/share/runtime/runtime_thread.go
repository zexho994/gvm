package runtime

// 映射到java中的一个thread todo
type Thread struct {
	pc    int
	stack *Stack
}

func (t *Thread) PC() int {
	return t.pc
}

func (t *Thread) SetPC(newPc int) {
	t.pc = newPc
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

func (t *Thread) IsStackEmpty() bool {
	return t.stack.isEmtpy()
}

func NewThread() *Thread {
	return &Thread{
		stack: NewStack(1024),
	}
}
