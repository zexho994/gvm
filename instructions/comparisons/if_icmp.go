package comparisons

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

// less than or equals
type If_ICMPLE struct {
	base.InstructionIndex16
}

type If_ICMPGE struct {
	base.InstructionIndex16
}

type If_ICMPEQ struct {
	base.InstructionIndex16
}

type If_ICMPNE struct {
	base.InstructionIndex16
}

type If_ICMPLT struct {
	base.InstructionIndex16
}

type If_ICMPGT struct {
	base.InstructionIndex16
}

type If_ACMPEQ struct {
	base.InstructionIndex16
}

type If_ACMPNE struct {
	base.InstructionIndex16
}

// to branch if and only if val1 less or equals val2
func (icmp *If_ICMPLE) Execute(frame *runtime.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	goNext := icmp.Index
	if val1 <= val2 {
		base.Branch(frame, int(goNext))
	}
}

// to branch if and only if val1 great or equals val2
func (icmp *If_ICMPGE) Execute(frame *runtime.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	goNext := icmp.Index
	if val1 >= val2 {
		base.Branch(frame, int(goNext))
	}
}

func (icmp *If_ICMPEQ) Execute(frame *runtime.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	if val1 == val2 {
		base.Branch(frame, int(icmp.Index))
	}
}

func (icmp *If_ICMPGT) Execute(frame *runtime.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	if val1 > val2 {
		base.Branch(frame, int(icmp.Index))
	}
}

func (acmp *If_ACMPNE) Execute(frame *runtime.Frame) {
	val2 := frame.OperandStack().PopRef()
	val1 := frame.OperandStack().PopRef()
	if val1 != val2 {
		base.Branch(frame, int(acmp.Index))
	}
}

func (icmp *If_ICMPNE) Execute(frame *runtime.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	if val1 != val2 {
		base.Branch(frame, int(icmp.Index))
	}
}

func (icmp If_ICMPLT) Execute(frame *runtime.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	if val1 < val2 {
		base.Branch(frame, int(icmp.Index))
	}
}

func (acmp *If_ACMPEQ) Execute(frame *runtime.Frame) {
	val2 := frame.OperandStack().PopRef()
	val1 := frame.OperandStack().PopRef()
	if val1 == val2 {
		base.Branch(frame, int(acmp.Index))
	}
}
