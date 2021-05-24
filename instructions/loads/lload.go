package loads

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type LLOAD struct {
	base.InstructionIndex16
}

type LLOAD_0 struct {
	base.InstructionIndex0
}

type LLOAD_1 struct {
	base.InstructionIndex0
}

type LLOAD_2 struct {
	base.InstructionIndex0
}

type LLOAD_3 struct {
	base.InstructionIndex0
}

/*
according index to load a var from frame.localVars
*/
func _lload(frame *runtime.Frame, index uint) {
	val := frame.GetLong(index)
	frame.OperandStack().PushLong(val)
}

/*
execute LLOAD
the index is stored inside the instrution
*/
func (self *LLOAD) Execute(frame *runtime.Frame) {
	_lload(frame, uint(self.Index))
}

/*
execute LLOAD_0
the index is zero
*/
func (self *LLOAD_0) Execute(frame *runtime.Frame) {
	_lload(frame, 0)
}

/*
see LLOAD_0's Execute
*/
func (self *LLOAD_1) Execute(frame *runtime.Frame) {
	_lload(frame, 1)
}

func (self *LLOAD_2) Execute(frame *runtime.Frame) {
	_lload(frame, 2)
}

func (self *LLOAD_3) Execute(frame *runtime.Frame) {
	_lload(frame, 3)
}
