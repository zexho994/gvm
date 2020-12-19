package instructions

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/instructions/control"
	"github.com/zouzhihao-994/gvm/src/share/instructions/references"
	"github.com/zouzhihao-994/gvm/src/share/instructions/stack"
)

var (
	dup           = &stack.Dup{}
	dup_x1        = &stack.Dup_X1{}
	dup_x2        = &stack.Dup_X2{}
	dup2          = &stack.Dup2{}
	dup2_x1       = &stack.Dup2_X1{}
	dup2_x2       = &stack.Dup2_X2{}
	voidReturn    = &control.RETURN{}           //177
	getStatic     = &references.GET_STATIC{}    // 178
	invokeStatic  = &references.INVOKE_STATIC{} // 184
	invokeSpecial = &references.INVOKE_SPECIAL{}
	_new          = &references.NEW{} //187
)

func NewInstruction(opcode byte) base.Base_Instruction {
	switch opcode {
	case 0x59:
		return dup
	case 0x5a:
		return dup_x1
	case 0x5b:
		return dup_x2
	case 0x5c:
		return dup2
	case 0x5d:
		return dup2_x1
	case 0x5e:
		return dup2_x2
	case 0xb1:
		return voidReturn
	case 0xb2:
		return getStatic
	case 0xb7:
		return invokeSpecial
	case 0xb8:
		return invokeStatic
	case 0xbb:
		return _new
	default:
		panic(fmt.Errorf("Unsupported opcode : 0x%x!", opcode))
	}
}
