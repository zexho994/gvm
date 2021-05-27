package math

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type Land struct {
	base.InstructionIndex0
}

func (l *Land) Execute(frame *runtime.Frame) {
	v2 := frame.PopLong()
	v1 := frame.PopLong()
	frame.PushLong(v1 & v2)
}
