package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

type ARRAY_LENGTH struct {
	base.InstructionIndex0
}

func (i *ARRAY_LENGTH) Execute(frame *runtime.Frame) {
	arrayRef := frame.PopRef()
	utils.AssertFalse(arrayRef == nil, "NullPointerException")
	arrayLength := arrayRef.ArrayLength()
	frame.PushInt(int32(arrayLength))
}
