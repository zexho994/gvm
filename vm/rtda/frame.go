package rtda

type Frame struct {
	// like the LinkedList's next point
	lower *Frame
	// Save the pointer of the local variable table corresponding to the stack frame
	localVars    LocalVars
	operandStack *OperandStack
}

/*
The value of maxLocals and maxStack can be calculated at compile time
can see the classfile.method_info's Code Attribute
*/
func newFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:  newLocalVars(maxLocals),
		operandSta: newOperandStack(maxStack),
	}
}
