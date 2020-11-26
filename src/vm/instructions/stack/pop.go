package stack

import "github.com/zouzhihao-994/gvm/src/vm/instructions/base"
import "github.com/zouzhihao-994/gvm/src/vm/runtime"

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

func (self *POP) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}
