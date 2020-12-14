package runtime

import "github.com/zouzhihao-994/gvm/src/share/jclass"

// 一个Frame对应着一个已调用而且未结束的方法
// TODO：栈的大小支持自动 扩/缩 , 如果扩到极限仍然发送内容不足的情况抛出 OutOfMemoryError 异常
type Frame struct {
	nextPc       int
	next         *Frame
	localVars    *LocalVars
	operandStack *OperandStack
	method       *jclass.MethodInfo
	thread       *Thread
}

func (f *Frame) SetNextPC(pc int) {
	f.nextPc = pc
}

func (f *Frame) NextPC() int {
	return f.nextPc
}

func (f *Frame) Method() *jclass.MethodInfo {
	return f.method
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

func NewFrame(maxlocals, maxStack uint16, method *jclass.MethodInfo, thread *Thread) *Frame {
	return &Frame{
		localVars:    NewLocalVars(maxlocals),
		operandStack: NewOperandStack(maxStack),
		method:       method,
		thread:       thread,
	}
}
