package runtime

// 一个Frame对应着一个已调用而且未结束的方法
// TODO：栈的大小支持自动 扩/缩 , 如果扩到极限仍然发送内容不足的情况抛出 OutOfMemoryError 异常
type Frame struct {
	nextPc       int
	next         *Frame
	localVars    *LocalVars
	operandStack *OperandStack
}

func (f *Frame) SetNextPC(pc int) {
	f.nextPc = pc
}

func (f *Frame) NextPC() int {
	return f.nextPc
}

func NewFrame(maxlocals, maxStack uint16) *Frame {
	return &Frame{
		localVars:    NewLocalVars(maxlocals),
		operandStack: NewOperandStack(maxStack),
	}
}
