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

func (self *AconstNull) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushRef(nil)
}

func (self *Dconst0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

func (self *Dconst1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

func (self *Fconst0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

func (self *Fconst1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

func (self *Fconst2) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

func (self *Iconst0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(0)
}

func (self *Iconst1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(1)
}

func (self *Iconst2) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(2)
}

func (self *Iconst3) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(3)
}

func (self *Iconst4) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(4)
}

func (self *Iconst5) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(5)
}

func (self *IconstM1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(-1)
}

func (self *Lconst0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushLong(0.0)
}

func (self *Lconst1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushLong(1.0)
}
