package references

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

type NEW_ARRAY struct {
	count    uint
	arrayref *jclass.JClass_Instance
	base.InstructionIndex8
}

func (i *NEW_ARRAY) Execute(frame *runtime.Frame) {
	arrayCount := frame.OperandStack().PopInt()
	exception.AssertFalse(arrayCount < 0, "NegativeArraySizeException")
	jclass.NewJarray(uint32(arrayCount), i.Index)
}
