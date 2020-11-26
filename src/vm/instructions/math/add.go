package math

import (
	"github.com/zouzhihao-994/gvm/src/vm/instructions/base"
	"github.com/zouzhihao-994/gvm/src/vm/runtime"
)

type IADD struct {
	base.NoOperandsInstruction
}

func (self *IADD) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	addResult := v1 + v2
	stack.PushInt(addResult)
}

type LADD struct {
	base.NoOperandsInstruction
}

func (self *LADD) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	addResult := v1 + v2
	stack.PushLong(addResult)
}

type FADD struct {
	base.NoOperandsInstruction
}

func (self *FADD) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	addResult := v1 + v2
	stack.PushFloat(addResult)
}

type DADD struct {
	base.NoOperandsInstruction
}

func (self *DADD) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	addResult := v1 + v2
	stack.PushDouble(addResult)
}
