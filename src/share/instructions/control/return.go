package control

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

type RETURN struct{ base.InstructionIndex0 } // Return void from method

type ARETURN struct{ base.InstructionIndex0 } // Return reference from met

type DRETURN struct{ base.InstructionIndex0 } // Return double from method

type FRETURN struct{ base.InstructionIndex0 } // Return float from method

type IRETURN struct{ base.InstructionIndex0 } // Return int from method

type LRETURN struct{ base.InstructionIndex0 } // Return long from method

/*
void返回类型的，直接返回顶部栈帧
*/
func (self *RETURN) Execute(frame *runtime.Frame) { frame.Thread().Pop() }

/*
int返回类型的，弹出当前栈帧的栈顶元素，push到最新的栈顶栈帧中（即调用方）
*/
func (self *IRETURN) Execute(frame *runtime.Frame) {
	thread := frame.Thread()
	currentFrame := thread.Pop()
	invokerFrame := thread.Peek()
	retVal := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retVal)
}

func (self *DRETURN) Execute(frame *runtime.Frame) {
	thread := frame.Thread()
	currentFrame := thread.Pop()
	invokerFrame := thread.Peek()
	retVal := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(retVal)
}

func (self *ARETURN) Execute(frame *runtime.Frame) {
	thread := frame.Thread()
	currentFrame := thread.Pop()
	invokerFrame := thread.Peek()
	retVal := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(retVal)
}

func (self *LRETURN) Execute(frame *runtime.Frame) {
	thread := frame.Thread()
	currentFrame := thread.Pop()
	invokerFrame := thread.Peek()
	retVal := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(retVal)
}

func (self *FRETURN) Execute(frame *runtime.Frame) {
	thread := frame.Thread()
	currentFrame := thread.Pop()
	invokerFrame := thread.Peek()
	retVal := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(retVal)
}
