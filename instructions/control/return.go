package control

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type RETURN struct{ base.InstructionIndex0 } // Return void from method

type ARETURN struct{ base.InstructionIndex0 } // Return reference from met

type DRETURN struct{ base.InstructionIndex0 } // Return double from method

type FRETURN struct{ base.InstructionIndex0 } // Return float from method

type IRETURN struct{ base.InstructionIndex0 } // Return int from method

type LRETURN struct{ base.InstructionIndex0 } // Return long from method

// Execute void返回类型的，直接返回顶部栈帧
func (r *RETURN) Execute(frame *runtime.Frame) { frame.Thread.PopFrame() }

// Execute int返回类型的，弹出当前栈帧的栈顶元素，push到最新的栈顶栈帧中（即调用方）
func (r *IRETURN) Execute(frame *runtime.Frame) {
	currentFrame := frame.Thread.PopFrame()
	invokerFrame := frame.Thread.PeekFrame()
	retVal := currentFrame.PopInt()
	invokerFrame.PushInt(retVal)
}

func (r *DRETURN) Execute(frame *runtime.Frame) {
	currentFrame := frame.PopFrame()
	invokerFrame := frame.PeekFrame()
	retVal := currentFrame.PopDouble()
	invokerFrame.PushDouble(retVal)
}

func (r *ARETURN) Execute(frame *runtime.Frame) {
	currentFrame := frame.PopFrame()
	invokerFrame := frame.PeekFrame()
	retVal := currentFrame.PopRef()
	invokerFrame.PushRef(retVal)
}

func (r *LRETURN) Execute(frame *runtime.Frame) {
	currentFrame := frame.PopFrame()
	invokerFrame := frame.PeekFrame()
	retVal := currentFrame.PopLong()
	invokerFrame.PushLong(retVal)
}

func (r *FRETURN) Execute(frame *runtime.Frame) {
	currentFrame := frame.PopFrame()
	invokerFrame := frame.PeekFrame()
	retVal := currentFrame.PopFloat()
	invokerFrame.PushFloat(retVal)
}
