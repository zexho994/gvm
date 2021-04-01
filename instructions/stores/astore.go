package stores

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

/*

 */
type ASTORE struct {
	base.InstructionIndex8
}

type ASTORE_0 struct {
	base.InstructionIndex0
}

type ASTORE_1 struct {
	base.InstructionIndex0
}

type ASTORE_2 struct {
	base.InstructionIndex0
}

type ASTORE_3 struct {
	base.InstructionIndex0
}

func _astore(frame *runtime.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}

func (self *ASTORE) Execute(frame *runtime.Frame) {
	_astore(frame, uint(self.Index))
}

func (self *ASTORE_0) Execute(frame *runtime.Frame) {
	_astore(frame, 0)
}

func (self *ASTORE_1) Execute(frame *runtime.Frame) {
	_astore(frame, 1)
}

func (self *ASTORE_2) Execute(frame *runtime.Frame) {
	_astore(frame, 2)
}

func (self *ASTORE_3) Execute(frame *runtime.Frame) {
	_astore(frame, 3)
}
