package rtda

type Thread struct {
	pc    int
	stack *Stack
}

/*
创建新的线程
线程的栈大小为1024字节
*/
func NewThread() *Thread {
	return &Thread{stack: newStack(1024)}
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return NewFrame(self, maxLocals, maxStack)
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}
