package stack

import "../base"
import "../../rtda"

/*
弹出一个操作数栈位大小
适用于int，float等
*/
type POP struct {
	base.NoOperandsInstruction
}

/*
弹出两个操作数栈位大小
例如double，long
*/
type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}
