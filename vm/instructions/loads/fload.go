package loads

import (
	"../../instructions/base"
	"../../rtda"
)

type FLOAD struct {
	base.Index16Instruction
}

type FLOAD_0 struct {
	base.NoOperandsInstruction
}

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

type FLOAD_2 struct {
	base.NoOperandsInstruction
}

type FLOAD_3 struct {
	base.NoOperandsInstruction
}

/*
according index to load a var from frame.localVars
*/
func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

/*
execute FLOAD
the index is stored inside the instrution
*/
func (self *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, self.Index)
}

/*
execute FLOAD_0
the index is zero
*/
func (self *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

/*
see FLOAD_0's Execute
*/
func (self *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

func (self *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

func (self *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
