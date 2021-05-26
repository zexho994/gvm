package conversions

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type I2f struct {
	base.InstructionIndex0
}

type I2l struct {
	base.InstructionIndex0
}

type F2i struct {
	base.InstructionIndex0
}

func (i I2f) Execute(frame *runtime.Frame) {
	frame.PushFloat(float32(frame.PopInt()))
}

func (i I2l) Execute(frame *runtime.Frame) {
	frame.PushLong(int64(frame.PopInt()))
}

func (i F2i) Execute(frame *runtime.Frame) {
	frame.PushInt(int32(frame.PopFloat()))
}
