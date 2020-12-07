package runtime

type Frame struct {
	next         *Frame
	localVars    *LocalVars
	operandStack *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    NewLocalVars(maxLocals),
		operandStack: NewOperandStack(maxStack),
	}
}
