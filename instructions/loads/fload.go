package loads

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type FLOAD struct {
	base.InstructionIndex8
}

type Fload0 struct {
	base.InstructionIndex0
}

type Fload1 struct {
	base.InstructionIndex0
}

type Fload2 struct {
	base.InstructionIndex0
}

type Fload3 struct {
	base.InstructionIndex0
}

/*
according index to load a var from frame.localVars
*/
func _fload(frame *runtime.Frame, index uint) {
	val := frame.GetFloat(index)
	frame.PushFloat(val)
}

// Execute execute FLOAD
// the index is stored inside the instrution
func (self *FLOAD) Execute(frame *runtime.Frame) {
	_fload(frame, uint(self.Index))
}

// Execute execute Fload0
// the index is zero
func (self *Fload0) Execute(frame *runtime.Frame) {
	_fload(frame, 0)
}

// Execute see Fload0's Execute
func (self *Fload1) Execute(frame *runtime.Frame) {
	_fload(frame, 1)
}

func (self *Fload2) Execute(frame *runtime.Frame) {
	_fload(frame, 2)
}

func (self *Fload3) Execute(frame *runtime.Frame) {
	_fload(frame, 3)
}
