package constants

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type AconstNull struct{ base.InstructionIndex0 }

type Dconst0 struct{ base.InstructionIndex0 }

type Dconst1 struct{ base.InstructionIndex0 }

type Fconst0 struct{ base.InstructionIndex0 }

type Fconst1 struct{ base.InstructionIndex0 }

type Fconst2 struct{ base.InstructionIndex0 }

type IconstM1 struct{ base.InstructionIndex0 }

type Iconst0 struct{ base.InstructionIndex0 }

type Iconst1 struct{ base.InstructionIndex0 }

type Iconst2 struct{ base.InstructionIndex0 }

type Iconst3 struct{ base.InstructionIndex0 }

type Iconst4 struct{ base.InstructionIndex0 }

type Iconst5 struct{ base.InstructionIndex0 }

type Lconst0 struct{ base.InstructionIndex0 }

type Lconst1 struct{ base.InstructionIndex0 }

func (*AconstNull) Execute(frame *runtime.Frame) {
	frame.PushRef(nil)
}

func (*Dconst0) Execute(frame *runtime.Frame) {
	frame.PushDouble(0.0)
}

func (*Dconst1) Execute(frame *runtime.Frame) {
	frame.PushDouble(1.0)
}

func (*Fconst0) Execute(frame *runtime.Frame) {
	frame.PushFloat(0.0)
}

func (*Fconst1) Execute(frame *runtime.Frame) {
	frame.PushFloat(1.0)
}

func (*Fconst2) Execute(frame *runtime.Frame) {
	frame.PushFloat(2.0)
}

func (*Iconst0) Execute(frame *runtime.Frame) {
	frame.PushInt(0)
}

func (*Iconst1) Execute(frame *runtime.Frame) {
	frame.PushInt(1)
}

func (*Iconst2) Execute(frame *runtime.Frame) {
	frame.PushInt(2)
}

func (*Iconst3) Execute(frame *runtime.Frame) {
	frame.PushInt(3)
}

func (*Iconst4) Execute(frame *runtime.Frame) {
	frame.PushInt(4)
}

func (*Iconst5) Execute(frame *runtime.Frame) {
	frame.PushInt(5)
}

func (*IconstM1) Execute(frame *runtime.Frame) {
	frame.PushInt(-1)
}

func (*Lconst0) Execute(frame *runtime.Frame) {
	frame.PushLong(0.0)
}

func (*Lconst1) Execute(frame *runtime.Frame) {
	frame.PushLong(1.0)
}
