package stores

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

// Store double into local variable
type DSTORE struct{ base.InstructionIndex16 }

func (self *DSTORE) Execute(frame *runtime.Frame) {
	_dstore(frame, uint(self.Index))
}

type DSTORE_0 struct {
	base.InstructionIndex0
}

func (self *DSTORE_0) Execute(frame *runtime.Frame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct {
	base.InstructionIndex0
}

func (self *DSTORE_1) Execute(frame *runtime.Frame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct {
	base.InstructionIndex0
}

func (self *DSTORE_2) Execute(frame *runtime.Frame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct {
	base.InstructionIndex0
}

func (self *DSTORE_3) Execute(frame *runtime.Frame) {
	_dstore(frame, 3)
}

func _dstore(frame *runtime.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}
