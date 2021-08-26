package loads

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type LLOAD struct {
	base.InstructionIndex8
}

type Lload0 struct {
	base.InstructionIndex0
}

type Lload1 struct {
	base.InstructionIndex0
}

type Lload2 struct {
	base.InstructionIndex0
}

type Lload3 struct {
	base.InstructionIndex0
}

/*
according index to load a var from frame.localVars
*/
func _lload(frame *runtime.Frame, index uint) {
	val := frame.GetLong(index)
	frame.PushLong(val)
}

// Execute execute LLOAD
// the index is stored inside the instrution
func (self *LLOAD) Execute(frame *runtime.Frame) {
	_lload(frame, uint(self.Index))
}

// Execute execute Lload0
// the index is zero
func (self *Lload0) Execute(frame *runtime.Frame) {
	_lload(frame, 0)
}

// Execute see Lload0's Execute
func (self *Lload1) Execute(frame *runtime.Frame) {
	_lload(frame, 1)
}

func (self *Lload2) Execute(frame *runtime.Frame) {
	_lload(frame, 2)
}

func (self *Lload3) Execute(frame *runtime.Frame) {
	_lload(frame, 3)
}
