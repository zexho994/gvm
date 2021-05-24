package runtime

// Thread 映射到java中的一个thread todo
type Thread struct {
	PC uint
	*Stack
}

// IsFinished 线程任务是否执行完成
func (t Thread) IsFinished() bool {
	if t.IsEmtpy() {
		return true
	}
	return false
}
