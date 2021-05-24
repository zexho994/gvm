package loads

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type FLOAD struct {
	base.InstructionIndex8
}

type FLOAD_0 struct {
	base.InstructionIndex0
}

type FLOAD_1 struct {
	base.InstructionIndex0
}

type FLOAD_2 struct {
	base.InstructionIndex0
}

type FLOAD_3 struct {
	base.InstructionIndex0
}

/*
according index to load a var from frame.localVars
*/
func _fload(frame *runtime.Frame, index uint) {
	val := frame.GetFloat(index)
	frame.PushFloat(val)
}

/*
execute FLOAD
the index is stored inside the instrution
*/
func (self *FLOAD) Execute(frame *runtime.Frame) {
	_fload(frame, uint(self.Index))
}

/*
execute FLOAD_0
the index is zero
*/
func (self *FLOAD_0) Execute(frame *runtime.Frame) {
	_fload(frame, 0)
}

/*
see FLOAD_0's Execute
*/
func (self *FLOAD_1) Execute(frame *runtime.Frame) {
	_fload(frame, 1)
}

func (self *FLOAD_2) Execute(frame *runtime.Frame) {
	_fload(frame, 2)
}

func (self *FLOAD_3) Execute(frame *runtime.Frame) {
	_fload(frame, 3)
}
