package stores

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// Store float into local variable
type FSTORE struct{ base.InstructionIndex8 }

func (self *FSTORE) Execute(frame *runtime.Frame) {
	_fstore(frame, uint(self.Index))
}

type FSTORE_0 struct {
	base.InstructionIndex0
}

func (self *FSTORE_0) Execute(frame *runtime.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct {
	base.InstructionIndex0
}

func (self *FSTORE_1) Execute(frame *runtime.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct {
	base.InstructionIndex0
}

func (self *FSTORE_2) Execute(frame *runtime.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct {
	base.InstructionIndex0
}

func (self *FSTORE_3) Execute(frame *runtime.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *runtime.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
