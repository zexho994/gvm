package references

import (
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type ARRAY_LENGTH struct {
	base.InstructionIndex0
}

func (i *ARRAY_LENGTH) Execute(frame *runtime.Frame) {
	arrayRef := frame.OperandStack().PopRef()
	exception.AssertFalse(arrayRef == nil, "NullPointerException")
	arrayLength := arrayRef.ArrayLength()
	frame.OperandStack().PushInt(int32(arrayLength))
}
