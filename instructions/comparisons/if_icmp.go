package comparisons

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

// IfIcmple less than or equals
type IfIcmple struct {
	base.InstructionIndex16
}

type IfIcmpge struct {
	base.InstructionIndex16
}

type IfIcmpeq struct {
	base.InstructionIndex16
}

type IfIcmpne struct {
	base.InstructionIndex16
}

type IfIcmplt struct {
	base.InstructionIndex16
}

type IfIcmpgt struct {
	base.InstructionIndex16
}

type IfAcmpeq struct {
	base.InstructionIndex16
}

type IfAcmpne struct {
	base.InstructionIndex16
}

// Execute to branch if and only if val1 less or equals val2
func (icmp *IfIcmple) Execute(frame *runtime.Frame) {
	val2 := frame.PopInt()
	val1 := frame.PopInt()
	goNext := icmp.Index
	if val1 <= val2 {
		base.Branch(frame, int(goNext))
	}
}

// Execute to branch if and only if val1 great or equals val2
func (icmp *IfIcmpge) Execute(frame *runtime.Frame) {
	val2 := frame.PopInt()
	val1 := frame.PopInt()
	goNext := icmp.Index
	if val1 >= val2 {
		base.Branch(frame, int(goNext))
	}
}

func (icmp *IfIcmpeq) Execute(frame *runtime.Frame) {
	val2 := frame.PopInt()
	val1 := frame.PopInt()
	if val1 == val2 {
		base.Branch(frame, int(icmp.Index))
	}
}

func (icmp *IfIcmpgt) Execute(frame *runtime.Frame) {
	val2 := frame.PopInt()
	val1 := frame.PopInt()
	if val1 > val2 {
		base.Branch(frame, int(icmp.Index))
	}
}

func (acmp *IfAcmpne) Execute(frame *runtime.Frame) {
	val2 := frame.PopRef()
	val1 := frame.PopRef()
	if val1 != val2 {
		base.Branch(frame, int(acmp.Index))
	}
}

func (icmp *IfIcmpne) Execute(frame *runtime.Frame) {
	val2 := frame.PopInt()
	val1 := frame.PopInt()
	if val1 != val2 {
		base.Branch(frame, int(icmp.Index))
	}
}

func (icmp IfIcmplt) Execute(frame *runtime.Frame) {
	val2 := frame.PopInt()
	val1 := frame.PopInt()
	if val1 < val2 {
		base.Branch(frame, int(icmp.Index))
	}
}

func (acmp *IfAcmpeq) Execute(frame *runtime.Frame) {
	val2 := frame.PopRef()
	val1 := frame.PopRef()
	if val1 == val2 {
		base.Branch(frame, int(acmp.Index))
	}
}
