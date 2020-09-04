package stores

import (
	"../../instructions/base"
	"../../rtda"
	"fmt"
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
	val := frame.OperandStack().PopInt()
	fmt.Printf("[gvm][PushInt] %v 存储到局部变量表[%v]中\n", val, index)
	frame.LocalVars().SetInt(index, val)
}

func (self *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, self.Index)
}

func (self *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

func (self *ISTORE_1) Execute(frame *rtda.Frame) {
	fmt.Println("[gvm][istore_1] 操作数栈存储数到局部变量表[1]中")
	_istore(frame, 1)
}

func (self *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

func (self *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}
