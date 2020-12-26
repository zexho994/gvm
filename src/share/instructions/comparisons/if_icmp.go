package comparisons

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// less than or equals
type If_ICMPLE struct {
	base.InstructionIndex16
}

func (icmp *If_ICMPLE) Execute(frame *runtime.Frame) {
	val1 := frame.OperandStack().PopInt()
	val2 := frame.OperandStack().PopInt()
	// true
	if val1 <= val2 {

	} else {
		// false
	}

}
