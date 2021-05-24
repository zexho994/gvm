package math

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type IINC struct {
	base.InstructionIndex16
}

func (i *IINC) Execute(frame *runtime.Frame) {
	idx := i.Index >> 8
	toAdd := i.Index & 0x0011
	old := frame.GetInt(uint(idx))
	frame.SetInt(uint(idx), int32(toAdd)+old)
}
