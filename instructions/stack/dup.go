package stack

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type Dup struct {
	base.InstructionIndex0
}

type DupX1 struct {
	base.InstructionIndex0
}

type DupX2 struct {
	base.InstructionIndex0
}

type Dup2 struct {
	base.InstructionIndex0
}

type Dup2X1 struct {
	base.InstructionIndex0
}

type Dup2X2 struct {
	base.InstructionIndex0
}

// Execute Duplicate the top operandStack value
func (d *Dup) Execute(frame *runtime.Frame) {
	stack := frame
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

// Execute Duplicate the top operand stack value and insert two values down
// before : top ->down 1,2,3,4
// after : 1,2,1,3,4 . top value 1 duplicate and then insert two values down
func (d *DupX1) Execute(frame *runtime.Frame) {
	f := frame
	slot1 := f.PopSlot()
	slot2 := f.PopSlot()
	f.PushSlot(slot1)
	f.PushSlot(slot2)
	f.PushSlot(slot1)
}

// Execute Duplicate the top operand stack value and insert three values down
// before : top ->down 1,2,3,4
// after : 1,2,3,1,4 . top value 1 duplicate and then insert three values down
func (d *DupX2) Execute(frame *runtime.Frame) {
	f := frame
	slot1 := f.PopSlot()
	slot2 := f.PopSlot()
	slot3 := f.PopSlot()
	f.PushSlot(slot1)
	f.PushSlot(slot3)
	f.PushSlot(slot2)
	f.PushSlot(slot1)
}

func (d *Dup2) Execute(frame *runtime.Frame) {
	stack := frame
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (d *Dup2X1) Execute(frame *runtime.Frame) {
	stack := frame
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (d *Dup2X2) Execute(frame *runtime.Frame) {
	stack := frame
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
