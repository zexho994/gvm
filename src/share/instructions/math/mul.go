package math

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

type IMUL struct {
	base.InstructionIndex0
}

type FMUL struct {
	base.InstructionIndex0
}

type DMUL struct {
	base.InstructionIndex0
}

type LMUL struct {
	base.InstructionIndex0
}

func (self IMUL) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}

func (self LMUL) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}

func (self DMUL) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}

func (self FMUL) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}
