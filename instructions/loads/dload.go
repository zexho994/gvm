package loads

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type DLOAD struct {
	base.InstructionIndex16
}

type Dload0 struct {
	base.InstructionIndex0
}

type Dload1 struct {
	base.InstructionIndex0
}

type Dload2 struct {
	base.InstructionIndex0
}

type Dload3 struct {
	base.InstructionIndex0
}

/*
according index to load a var from frame.localVars
*/
func _dload(frame *runtime.Frame, index uint) {
	val := frame.GetDouble(index)
	frame.PushDouble(val)
}

// Execute DLOAD
// the index is stored inside the instrution
func (self *DLOAD) Execute(frame *runtime.Frame) {
	_dload(frame, uint(self.Index))
}

// Execute execute Dload0
// the index is zero
func (self *Dload0) Execute(frame *runtime.Frame) {
	_dload(frame, 0)
}

// Execute see Dload0's Execute
func (self *Dload1) Execute(frame *runtime.Frame) {
	_dload(frame, 1)
}

func (self *Dload2) Execute(frame *runtime.Frame) {
	_dload(frame, 2)
}

func (self *Dload3) Execute(frame *runtime.Frame) {
	_dload(frame, 3)
}
