package stores

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

/*
加载指令，将操作数栈的指保存到局部变了表中
*/
type LSTORE struct {
	base.InstructionIndex16
}

type LSTORE_0 struct {
	base.InstructionIndex0
}

type LSTORE_1 struct {
	base.InstructionIndex0
}

type LSTORE_2 struct {
	base.InstructionIndex0
}

type LSTORE_3 struct {
	base.InstructionIndex0
}

func _lstore(frame *runtime.Frame, index uint) {
	val := frame.PopLong()
	frame.SetLong(index, val)
}

func (self *LSTORE) Execute(frame *runtime.Frame) {
	_lstore(frame, uint(self.Index))
}

func (self *LSTORE_0) Execute(frame *runtime.Frame) {
	_lstore(frame, 0)
}

func (self *LSTORE_1) Execute(frame *runtime.Frame) {
	_lstore(frame, 1)
}

func (self *LSTORE_2) Execute(frame *runtime.Frame) {
	_lstore(frame, 2)
}

func (self *LSTORE_3) Execute(frame *runtime.Frame) {
	_lstore(frame, 3)
}
