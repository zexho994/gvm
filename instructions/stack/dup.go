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
	slot := frame.PopSlot()
	frame.PushSlot(slot)
	frame.PushSlot(slot)
}

// Execute Duplicate the top operand stack value and insert two values down
// before : top ->down 1,2,3,4
// after : 1,2,1,3,4 . top value 1 duplicate and then insert two values down
func (d *DupX1) Execute(frame *runtime.Frame) {
	slot1 := frame.PopSlot()
	slot2 := frame.PopSlot()
	frame.PushSlot(slot1)
	frame.PushSlot(slot2)
	frame.PushSlot(slot1)
}

// Execute Duplicate the top operand stack value and insert three values down
// before : top ->down 1,2,3,4
// after : 1,2,3,1,4 . top value 1 duplicate and then insert three values down
func (d *DupX2) Execute(frame *runtime.Frame) {
	slot1 := frame.PopSlot()
	slot2 := frame.PopSlot()
	slot3 := frame.PopSlot()
	frame.PushSlot(slot1)
	frame.PushSlot(slot3)
	frame.PushSlot(slot2)
	frame.PushSlot(slot1)
}

func (d *Dup2) Execute(frame *runtime.Frame) {
	slot1 := frame.PopSlot()
	slot2 := frame.PopSlot()
	frame.PushSlot(slot2)
	frame.PushSlot(slot1)
	frame.PushSlot(slot2)
	frame.PushSlot(slot1)
}

func (d *Dup2X1) Execute(frame *runtime.Frame) {
	slot1 := frame.PopSlot()
	slot2 := frame.PopSlot()
	slot3 := frame.PopSlot()

	frame.PushSlot(slot2)
	frame.PushSlot(slot1)
	frame.PushSlot(slot3)
	frame.PushSlot(slot2)
	frame.PushSlot(slot1)
}

func (d *Dup2X2) Execute(frame *runtime.Frame) {
	slot1 := frame.PopSlot()
	slot2 := frame.PopSlot()
	slot3 := frame.PopSlot()
	slot4 := frame.PopSlot()

	frame.PushSlot(slot2)
	frame.PushSlot(slot1)
	frame.PushSlot(slot4)
	frame.PushSlot(slot3)
	frame.PushSlot(slot2)
	frame.PushSlot(slot1)
}
