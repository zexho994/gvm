package stores

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type ASTORE struct {
	base.InstructionIndex8
}

type Astore0 struct {
	base.InstructionIndex0
}

type Astore1 struct {
	base.InstructionIndex0
}

type Astore2 struct {
	base.InstructionIndex0
}

type Astore3 struct {
	base.InstructionIndex0
}

func _astore(frame *runtime.Frame, index uint) {
	val := frame.PopRef()
	frame.SetRef(index, val)
}

func (self *ASTORE) Execute(frame *runtime.Frame) {
	_astore(frame, uint(self.Index))
}

func (self *Astore0) Execute(frame *runtime.Frame) {
	_astore(frame, 0)
}

func (self *Astore1) Execute(frame *runtime.Frame) {
	_astore(frame, 1)
}

func (self *Astore2) Execute(frame *runtime.Frame) {
	_astore(frame, 2)
}

func (self *Astore3) Execute(frame *runtime.Frame) {
	_astore(frame, 3)
}
