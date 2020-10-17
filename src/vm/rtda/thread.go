package rtda

import "github.com/zouzhihao-994/gvm/src/vm/rtda/heap"

/*
线程
*/
type Thread struct {
	/*
		pc指针
		表示指令执行的位置
	*/
	pc int

	/*
		虚拟机栈属于线程的一个私有字段
		说明虚拟机栈是线程私有的
	*/
	stack *Stack
}

/*
创建新的线程
线程的栈大小为1024字节
*/
func NewThread() *Thread {
	return &Thread{stack: newStack(1024)}
}

/*
线程的虚拟机栈压入栈桢
*/
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

/*
虚拟机栈弹出栈桢
*/
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

/*
虚拟机顶的栈帧
*/
func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}

/*
获取当前虚拟机栈栈顶的栈桢
*/
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}
