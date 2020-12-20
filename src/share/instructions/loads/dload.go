package loads

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

type DLOAD struct {
	base.InstructionIndex16
}

type DLOAD_0 struct {
	base.InstructionIndex0
}

type DLOAD_1 struct {
	base.InstructionIndex0
}

type DLOAD_2 struct {
	base.InstructionIndex0
}

type DLOAD_3 struct {
	base.InstructionIndex0
}

/*
according index to load a var from frame.localVars
*/
func _dload(frame *runtime.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

/*
execute DLOAD
the index is stored inside the instrution
*/
func (self *DLOAD) Execute(frame *runtime.Frame) {
	_dload(frame, uint(self.Index))
}

/*
execute DLOAD_0
the index is zero
*/
func (self *DLOAD_0) Execute(frame *runtime.Frame) {
	_dload(frame, 0)
}

/*
see DLOAD_0's Execute
*/
func (self *DLOAD_1) Execute(frame *runtime.Frame) {
	_dload(frame, 1)
}

func (self *DLOAD_2) Execute(frame *runtime.Frame) {
	_dload(frame, 2)
}

func (self *DLOAD_3) Execute(frame *runtime.Frame) {
	_dload(frame, 3)
}
