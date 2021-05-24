package loads

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type ILOAD struct {
	base.InstructionIndex8
}

type Iload0 struct {
	base.InstructionIndex0
}

type Iload1 struct {
	base.InstructionIndex0
}

type Iload2 struct {
	base.InstructionIndex0
}

type Iload3 struct {
	base.InstructionIndex0
}

/*
according index to load a var from frame.localVars
*/
func _iload(frame *runtime.Frame, index uint) {
	val := frame.GetInt(index)
	frame.PushInt(val)
}

// Execute execute ILOAD
// the index is stored inside the instrution
func (self *ILOAD) Execute(frame *runtime.Frame) {
	_iload(frame, uint(self.Index))
}

// Execute execute Iload0
// the index is zero
func (self *Iload0) Execute(frame *runtime.Frame) {
	_iload(frame, 0)
}

// Execute see Iload0's Execute
func (self *Iload1) Execute(frame *runtime.Frame) {
	_iload(frame, 1)
}

func (self *Iload2) Execute(frame *runtime.Frame) {
	_iload(frame, 2)
}

func (self *Iload3) Execute(frame *runtime.Frame) {
	_iload(frame, 3)
}
