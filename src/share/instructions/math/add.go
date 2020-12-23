package math

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

type IADD struct {
	base.InstructionIndex0
}

type LADD struct {
	base.InstructionIndex0
}

type FADD struct {
	base.InstructionIndex0
}

type DADD struct {
	base.InstructionIndex0
}

func (this *IADD) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	addResult := v1 + v2
	stack.PushInt(addResult)
}

func (this *LADD) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	addResult := v1 + v2
	stack.PushLong(addResult)
}

func (this *FADD) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	addResult := v1 + v2
	stack.PushFloat(addResult)
}

func (this *DADD) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	addResult := v1 + v2
	stack.PushDouble(addResult)
}
