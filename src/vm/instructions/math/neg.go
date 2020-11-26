package math

import (
	"github.com/zouzhihao-994/gvm/src/vm/instructions/base"
	"github.com/zouzhihao-994/gvm/src/vm/runtime"
)

type INEG struct {
	base.NoOperandsInstruction
}

type LNEG struct {
	base.NoOperandsInstruction
}

type FNEG struct {
	base.NoOperandsInstruction
}

type DNEG struct {
	base.NoOperandsInstruction
}

func (self FNEG) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	n := -v1
	stack.PushFloat(n)
}

func (self INEG) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	n := -v1
	stack.PushInt(n)
}

func (self DNEG) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	n := -v1
	stack.PushDouble(n)
}

func (self LNEG) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	n := -v1
	stack.PushLong(n)
}
