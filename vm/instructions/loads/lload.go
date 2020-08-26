package loads

import (
	"../../instructions/base"
	"../../rtda"
)

type LLOAD struct {
	base.Index16Instruction
}

type LLOAD_0 struct {
	base.NoOperandsInstruction
}

type LLOAD_1 struct {
	base.NoOperandsInstruction
}

type LLOAD_2 struct {
	base.NoOperandsInstruction
}

type LLOAD_3 struct {
	base.NoOperandsInstruction
}

/*
according index to load a var from frame.localVars
*/
func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

/*
execute LLOAD
the index is stored inside the instrution
*/
func (self *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, uint(self.Index))
}

/*
execute LLOAD_0
the index is zero
*/
func (self *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

/*
see LLOAD_0's Execute
*/
func (self *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

func (self *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

func (self *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
