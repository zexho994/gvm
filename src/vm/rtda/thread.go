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
func (thread *Thread) PushFrame(frame *Frame) {
	thread.stack.push(frame)
}

/*
虚拟机栈弹出栈桢
*/
func (thread *Thread) PopFrame() *Frame {
	return thread.stack.pop()
}

/*
虚拟机顶的栈帧
*/
func (thread *Thread) TopFrame() *Frame {
	return thread.stack.top()
}

/*
获取当前虚拟机栈栈顶的栈桢
*/
func (thread *Thread) CurrentFrame() *Frame {
	return thread.stack.top()
}

func (thread *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(thread, method)
}

func (thread *Thread) PC() int {
	return thread.pc
}

func (thread *Thread) SetPC(pc int) {
	thread.pc = pc
}

func (thread *Thread) IsStackEmpty() bool {
	return thread.stack.isEmpty()
}
