package loads

import (
	"../../instructions/base"
	"../../rtda"
)

type DLOAD struct {
	base.Index16Instruction
}

type DLOAD_0 struct {
	base.NoOperandsInstruction
}

type DLOAD_1 struct {
	base.NoOperandsInstruction
}

type DLOAD_2 struct {
	base.NoOperandsInstruction
}

type DLOAD_3 struct {
	base.NoOperandsInstruction
}

/*
according index to load a var from frame.localVars
*/
func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

/*
execute DLOAD
the index is stored inside the instrution
*/
func (self *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, self.Index)
}

/*
execute DLOAD_0
the index is zero
*/
func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

/*
see DLOAD_0's Execute
*/
func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}
