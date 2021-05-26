package math

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type Lshl struct {
	base.InstructionIndex0
}

func (l *Lshl) Execute(frame *runtime.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopLong()
	frame.PushLong(v1 << (v2 & 0x3f))
}
