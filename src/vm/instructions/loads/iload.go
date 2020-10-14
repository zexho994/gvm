package loads

import "../base"
import "../../rtda"

type ILOAD struct {
	base.Index8Instruction
}

type ILOAD_0 struct {
	base.NoOperandsInstruction
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

/*
according index to load a var from frame.localVars
*/
func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

/*
execute ILOAD
the index is stored inside the instrution
*/
func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, self.Index)
}

/*
execute ILOAD_0
the index is zero
*/
func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

/*
see ILOAD_0's Execute
*/
func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
