package stack

import "../base"
import "../../rtda"

type Dup struct {
	base.NoOperandsInstruction
}

type Dup_X1 struct {
	base.NoOperandsInstruction
}

type Dup_X2 struct {
	base.NoOperandsInstruction
}

type Dup2 struct {
	base.NoOperandsInstruction
}

type Dup2_X1 struct {
	base.NoOperandsInstruction
}

type Dup2_X2 struct {
	base.NoOperandsInstruction
}

// Duplicate the top operandStack value
func (self *Dup) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

/*
Duplicate the top operand stack value and insert two values down
before : top ->down 1,2,3,4
after : 1,2,1,3,4 . top value 1 duplicate and then insert two values down
*/
func (self *Dup_X1) Execute(frame *rtda.Frame) {
	f := frame.OperandStack()
	slot1 := f.PopSlot()
	slot2 := f.PopSlot()
	f.PushSlot(slot1)
	f.PushSlot(slot2)
	f.PushSlot(slot1)
}

/*
Duplicate the top operand stack value and insert three values down
before : top ->down 1,2,3,4
after : 1,2,3,1,4 . top value 1 duplicate and then insert three values down
*/
func (self *Dup_X2) Execute(frame *rtda.Frame) {
	f := frame.OperandStack()
	slot1 := f.PopSlot()
	slot2 := f.PopSlot()
	slot3 := f.PopSlot()
	f.PushSlot(slot1)
	f.PushSlot(slot3)
	f.PushSlot(slot2)
	f.PushSlot(slot1)
}

func (self *Dup2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (self *Dup2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (self *Dup2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
