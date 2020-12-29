package stores

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// pop ref from stack and store the ref to localvars
type IASTORE struct {
	base.InstructionIndex0
}

func (i *IASTORE) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	idx := frame.OperandStack().PopInt()
	array := frame.OperandStack().PopRef()
	exception.AssertFalse(array == nil, "NullPointerException")
	array.ArrayData().SetIVal(idx, val)
}
