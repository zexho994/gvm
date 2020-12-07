package runtime

// 一个Frame对应着一个已调用而且未结束的方法
// TODO：栈的大小支持自动 扩/缩 , 如果扩到极限仍然发送内容不足的情况抛出 OutOfMemoryError 异常
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
