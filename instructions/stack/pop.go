package stack

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

/*
弹出一个操作数栈位大小
适用于int，float等
*/
type POP struct {
	base.InstructionIndex0
}

/*
弹出两个操作数栈位大小
例如double，long
*/
type POP2 struct {
	base.InstructionIndex0
}

func (self *POP) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (self *POP2) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
