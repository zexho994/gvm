package loads

import (
	"github.com/zouzhihao-994/gvm/src/vm/instructions/base"
	"github.com/zouzhihao-994/gvm/src/vm/runtime"
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
func _fload(frame *runtime.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

/*
execute FLOAD
the index is stored inside the instrution
*/
func (self *FLOAD) Execute(frame *runtime.Frame) {
	_fload(frame, self.Index)
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
