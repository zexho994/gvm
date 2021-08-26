package stores

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

// DSTORE Store double into local variable
type DSTORE struct{ base.InstructionIndex8 }

func (self *DSTORE) Execute(frame *runtime.Frame) {
	_dstore(frame, uint(self.Index))
}

type Dstore0 struct {
	base.InstructionIndex0
}

func (self *Dstore0) Execute(frame *runtime.Frame) {
	_dstore(frame, 0)
}

type Dstore1 struct {
	base.InstructionIndex0
}

func (self *Dstore1) Execute(frame *runtime.Frame) {
	_dstore(frame, 1)
}

type Dstore2 struct {
	base.InstructionIndex0
}

func (self *Dstore2) Execute(frame *runtime.Frame) {
	_dstore(frame, 2)
}

type Dstore3 struct {
	base.InstructionIndex0
}

func (self *Dstore3) Execute(frame *runtime.Frame) {
	_dstore(frame, 3)
}

func _dstore(frame *runtime.Frame, index uint) {
	val := frame.PopDouble()
	frame.SetDouble(index, val)
}
