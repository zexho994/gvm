package stores

import (
	"../../instructions/base"
	"../../rtda"
)

/*
加载指令，将操作数栈的指保存到局部变了表中
*/
type ISTORE struct {
	base.Index8Instruction
}

type ISTORE_0 struct {
	base.NoOperandsInstruction
}

type ISTORE_1 struct {
	base.NoOperandsInstruction
}

type ISTORE_2 struct {
	base.NoOperandsInstruction
}

type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (self *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, self.Index)
}

func (self *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

func (self *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

func (self *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

func (self *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}
