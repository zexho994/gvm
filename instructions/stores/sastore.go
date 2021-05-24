package stores

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

// pop ref from stack and store the ref to localvars
type SASTORE struct {
	base.InstructionIndex0
}

func (i *SASTORE) Execute(frame *runtime.Frame) {
	val := frame.PopInt()
	idx := frame.PopInt()
	array := frame.PopRef()
	utils.AssertFalse(array == nil, "NullPointerException")
	array.ArrayData().SetCVal(idx, int8(val))
}
