package stack

import "../../instructions/base"
import "../../rtda"

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
