package stack

import "github.com/zouzhihao-994/gvm/src/vm/instructions/base"
import "github.com/zouzhihao-994/gvm/src/vm/rtda"

type SWAP struct {
	base.NoOperandsInstruction
}

/*
Swap the two values at the top of the stack
*/
func (self *SWAP) Execute(frmae *rtda.Frame) {
	stack := frmae.OperandStack()
	stack1 := stack.PopSlot()
	stack2 := stack.PopSlot()
	stack.PushSlot(stack1)
	stack.PushSlot(stack2)
}
