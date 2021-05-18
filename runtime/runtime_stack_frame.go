package runtime

import (
	"github.com/zouzhihao-994/gvm/klass"
)

// Frame 一个Frame对应着一个已调用而且未结束的方法
// TODO：栈的大小支持自动 扩/缩 , 如果扩到极限仍然发送内容不足的情况抛出 OutOfMemoryError 异常
type Frame struct {
	pc           uint
	next         *Frame
	localVars    *LocalVars
	operandStack *OperandStack
	method       *klass.MethodInfo
	thread       *Thread
}

func (f *Frame) LocalVars() *LocalVars {
	return f.localVars
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) SetPC(pc uint) {
	f.pc = pc
}

func (f *Frame) PC() uint {
	return f.pc
}

func (f *Frame) Method() *klass.MethodInfo {
	return f.method
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

// RevertPC 重置帧指针
// 在执行 inst.Execute() 方法之前会将 frame 的 pc 指针后移
// 而在某些 Execute() 方法中，发送类例如 pushFrame() 操作，为了保证新加入frame会在下次执行
// 就将frame的指针重置为thread的pc，
// 选择重置为thread.pc的而不是简单的进行pc--，因为除了获取操作码会进行pc++,在读取操作数的时候也会进行不同长度的pc++
func (f *Frame) RevertPC() {
	f.pc = f.thread.PC
}

func NewFrame(maxlocals, maxStack uint16, method *klass.MethodInfo, thread *Thread) *Frame {
	return &Frame{
		localVars:    NewLocalVars(maxlocals),
		operandStack: NewOperandStack(maxStack),
		method:       method,
		thread:       thread,
	}
}
