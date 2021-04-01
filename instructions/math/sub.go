package math

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type ISUB struct {
	base.InstructionIndex0
}

type LSUB struct {
	base.InstructionIndex0
}

type DSUB struct {
	base.InstructionIndex0
}

type FSUB struct {
	base.InstructionIndex0
}

func (sub *ISUB) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	subResult := v2 - v1
	stack.PushInt(subResult)
}

func (sub *LSUB) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	subResule := v2 - v1
	stack.PushLong(subResule)
}

func (sub *DSUB) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	subResult := v2 - v1
	stack.PushDouble(subResult)
}

func (sub *FSUB) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	subResult := v2 - v1
	stack.PushFloat(subResult)
}
