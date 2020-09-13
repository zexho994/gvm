package rtda

import "./heap"

type Frame struct {
	// like the LinkedList's next point
	lower *Frame
	// Save the pointer of the local variable table corresponding to the stack frame
	localVars    *LocalVars
	operandStack *OperandStack
	method       *heap.Method
	// 栈桢结构里
	thread *Thread
	// 下一个指令
	nextPc int
}

/*
The value of maxLocals and maxStack can be calculated at compile time
can see the classfile.method_info's Code Attribute
*/
func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    NewLocalVars(method.MaxLocals()),
		operandStack: NewOperandStack(method.MaxStack()),
	}
}

func (self Frame) Method() *heap.Method {
	return self.method
}

func (self Frame) LocalVars() *LocalVars {
	return self.localVars
}

func (self Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) SetNextPC(next int) {
	self.nextPc = next
}

func (self Frame) NextPC() int {
	return self.nextPc
}
