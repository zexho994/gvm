package loads

import (
	"github.com/zouzhihao-994/gvm/src/vm/instructions/base"
	"github.com/zouzhihao-994/gvm/src/vm/runtime"
)

type ALOAD struct {
	base.Index8Instruction
}

type ALOAD_0 struct {
	base.NoOperandsInstruction
}

type ALOAD_1 struct {
	base.NoOperandsInstruction
}

type ALOAD_2 struct {
	base.NoOperandsInstruction
}

type ALOAD_3 struct {
	base.NoOperandsInstruction
}

/*
according index to load a var from frame.localVars
*/
func _aload(frame *runtime.Frame, index uint) {
	rel := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(rel)
}

/*
execute ALOAD
the index is stored inside the instrution
*/
func (self *ALOAD) Execute(frame *runtime.Frame) {
	_aload(frame, self.Index)
}

/*
execute ALOAD_0
the index is zero
*/
func (self *ALOAD_0) Execute(frame *runtime.Frame) {
	_aload(frame, 0)
}

/*
see ALOAD_0's Execute
*/
func (self *ALOAD_1) Execute(frame *runtime.Frame) {
	_aload(frame, 1)
}

func (self *ALOAD_2) Execute(frame *runtime.Frame) {
	_aload(frame, 2)
}

func (self *ALOAD_3) Execute(frame *runtime.Frame) {
	_aload(frame, 3)
}
