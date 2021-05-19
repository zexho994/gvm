package loads

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type ALOAD struct {
	base.InstructionIndex8
}

type Aload0 struct {
	base.InstructionIndex0
}

type Aload1 struct {
	base.InstructionIndex0
}

type Aload2 struct {
	base.InstructionIndex0
}

type Aload3 struct {
	base.InstructionIndex0
}

// according index to load a var from frame.localVars
func _aload(frame *runtime.Frame, index uint) {
	rel := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(rel)
}

// Execute execute ALOAD
// the index is stored inside the instrution
func (a *ALOAD) Execute(frame *runtime.Frame) {
	_aload(frame, uint(a.Index))
}

// Execute execute Aload0
// the index is zero
func (a *Aload0) Execute(frame *runtime.Frame) {
	_aload(frame, 0)
}

// Execute see Aload0's
func (a *Aload1) Execute(frame *runtime.Frame) {
	_aload(frame, 1)
}

func (a *Aload2) Execute(frame *runtime.Frame) {
	_aload(frame, 2)
}

func (a *Aload3) Execute(frame *runtime.Frame) {
	_aload(frame, 3)
}
