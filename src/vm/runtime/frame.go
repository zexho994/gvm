package runtime

import "github.com/zouzhihao-994/gvm/src/vm/oops"

type Frame struct {
	// like the LinkedList's next point
	lower *Frame
	// Save the pointer of the local variable table corresponding to the stack frame
	localVars    *LocalVars
	operandStack *OperandStack
	method       *oops.Method
	// 栈桢结构里
	thread *Thread
	// 下一个指令
	nextPc int
}

/*
The value of maxLocals and maxStack can be calculated at compile time
can see the classfile.method_info's AttrCode Attribute
*/
func newFrame(thread *Thread, method *oops.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    NewLocalVars(method.MaxLocals()),
		operandStack: NewOperandStack(method.MaxStack()),
	}

}

func (frame Frame) Method() *oops.Method {
	return frame.method
}

func (frame Frame) LocalVars() *LocalVars {
	return frame.localVars
}

func (frame Frame) OperandStack() *OperandStack {
	return frame.operandStack
}

func (frame Frame) Thread() *Thread {
	return frame.thread
}

func (frame *Frame) SetNextPC(next int) {
	frame.nextPc = next
}

func (frame Frame) NextPC() int {
	return frame.nextPc
}

/*
重置pc指针
*/
func (frame *Frame) RevertNextPC() { frame.nextPc = frame.thread.pc }
