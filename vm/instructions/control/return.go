package control

import "../../instructions/base"
import "../../rtda"

type RETURN struct{ base.NoOperandsInstruction } // Return void from method

type ARETURN struct{ base.NoOperandsInstruction } // Return reference from met

type DRETURN struct{ base.NoOperandsInstruction } // Return double from method

type FRETURN struct{ base.NoOperandsInstruction } // Return float from method

type IRETURN struct{ base.NoOperandsInstruction } // Return int from method

type LRETURN struct{ base.NoOperandsInstruction } // Return long from method

/*
void返回类型的，直接返回顶部栈帧
*/
func (self *RETURN) Execute(frame *rtda.Frame) { frame.Thread().PopFrame() }

/*
int返回类型的，弹出当前栈帧的栈顶元素，push到最新的栈顶栈帧中（即调用方）
*/
func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retVal)
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(retVal)
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(retVal)
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(retVal)
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(retVal)
}
