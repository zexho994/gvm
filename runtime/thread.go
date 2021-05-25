package runtime

// Thread 映射到java中的一个thread
type Thread struct {
	pc uint
	*Stack
}

func (t Thread) ThreadPC() uint {
	return t.pc
}

func (t *Thread) SetThreadPC(pc uint) {
	t.pc = pc
}

func (t *Thread) RevertFramePC() {
	t.PeekFrame().RevertPC()
}
