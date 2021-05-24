package loads

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type ILOAD struct {
	base.InstructionIndex8
}

type ILOAD_0 struct {
	base.InstructionIndex0
}

type ILOAD_1 struct {
	base.InstructionIndex0
}

type ILOAD_2 struct {
	base.InstructionIndex0
}

type ILOAD_3 struct {
	base.InstructionIndex0
}

/*
according index to load a var from frame.localVars
*/
func _iload(frame *runtime.Frame, index uint) {
	val := frame.GetInt(index)
	frame.PushInt(val)
}

/*
execute ILOAD
the index is stored inside the instrution
*/
func (self *ILOAD) Execute(frame *runtime.Frame) {
	_iload(frame, uint(self.Index))
}

/*
execute ILOAD_0
the index is zero
*/
func (self *ILOAD_0) Execute(frame *runtime.Frame) {
	_iload(frame, 0)
}

/*
see ILOAD_0's Execute
*/
func (self *ILOAD_1) Execute(frame *runtime.Frame) {
	_iload(frame, 1)
}

func (self *ILOAD_2) Execute(frame *runtime.Frame) {
	_iload(frame, 2)
}

func (self *ILOAD_3) Execute(frame *runtime.Frame) {
	_iload(frame, 3)
}
