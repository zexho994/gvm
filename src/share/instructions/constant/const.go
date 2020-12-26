package constants

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

type ACONST_NULL struct{ base.InstructionIndex0 }

type DCONST_0 struct{ base.InstructionIndex0 }

type DCONST_1 struct{ base.InstructionIndex0 }

type FCONST_0 struct{ base.InstructionIndex0 }

type FCONST_1 struct{ base.InstructionIndex0 }

type FCONST_2 struct{ base.InstructionIndex0 }

type ICONST_M1 struct{ base.InstructionIndex0 }

type ICONST_0 struct{ base.InstructionIndex0 }

type ICONST_1 struct{ base.InstructionIndex0 }

type ICONST_2 struct{ base.InstructionIndex0 }

type ICONST_3 struct{ base.InstructionIndex0 }

type ICONST_4 struct{ base.InstructionIndex0 }

type ICONST_5 struct{ base.InstructionIndex0 }

type LCONST_0 struct{ base.InstructionIndex0 }

type LCONST_1 struct{ base.InstructionIndex0 }

func (self *ACONST_NULL) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushRef(nil)
}

func (self *DCONST_0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

func (self *DCONST_1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

func (self *FCONST_0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

func (self *FCONST_1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

func (self *FCONST_2) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

func (self *ICONST_0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(0)
}

func (self *ICONST_1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(1)
}

func (self *ICONST_2) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(2)
}

func (self *ICONST_3) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(3)
}

func (self *ICONST_4) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(4)
}

func (self *ICONST_5) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(5)
}

func (self *ICONST_M1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(-1)
}

func (self *LCONST_0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushLong(0.0)
}

func (self *LCONST_1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushLong(1.0)
}
