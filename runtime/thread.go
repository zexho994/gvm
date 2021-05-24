package runtime

// Thread 映射到java中的一个thread todo
type Thread struct {
	pc uint
	*Stack
}

// IsFinished 线程任务是否执行完成
func (t Thread) IsFinished() bool {
	if t.IsEmtpy() {
		return true
	}
	return false
}

func (t Thread) ThreadPC() uint {
	return t.pc
}

func (t *Thread) SetThradPC(pc uint) {
	t.pc = pc
}
