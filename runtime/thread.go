package runtime

// Thread 映射到java中的一个thread todo
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
