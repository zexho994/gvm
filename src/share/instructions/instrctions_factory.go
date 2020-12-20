package instructions

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/instructions/control"
	"github.com/zouzhihao-994/gvm/src/share/instructions/loads"
	"github.com/zouzhihao-994/gvm/src/share/instructions/references"
	"github.com/zouzhihao-994/gvm/src/share/instructions/stack"
)

var (
	iload         = &loads.ILOAD{}
	lload         = &loads.LLOAD{}
	fload         = &loads.FLOAD{}
	dload         = &loads.DLOAD{}
	aload         = &loads.ALOAD{}
	iload_0       = &loads.ILOAD_0{}
	iload_1       = &loads.ILOAD_1{}
	iload_2       = &loads.ILOAD_2{}
	iload_3       = &loads.ILOAD_3{}
	lload_0       = &loads.LLOAD_0{}
	lload_1       = &loads.LLOAD_1{}
	lload_2       = &loads.LLOAD_2{}
	lload_3       = &loads.LLOAD_3{}
	fload_0       = &loads.FLOAD_0{}
	fload_1       = &loads.FLOAD_1{}
	fload_2       = &loads.FLOAD_2{}
	fload_3       = &loads.FLOAD_3{}
	dload_0       = &loads.DLOAD_0{}
	dload_1       = &loads.DLOAD_1{}
	dload_2       = &loads.DLOAD_2{}
	dload_3       = &loads.DLOAD_3{}
	aload_0       = &loads.ALOAD_0{}
	aload_1       = &loads.ALOAD_1{}
	aload_2       = &loads.ALOAD_2{}
	aload_3       = &loads.ALOAD_3{}
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
	case 0x15:
		return iload
	case 0x16:
		return lload
	case 0x17:
		return fload
	case 0x18:
		return dload
	case 0x19:
		return aload
	case 0x1a:
		return iload_0
	case 0x1b:
		return iload_1
	case 0x1c:
		return iload_2
	case 0x1d:
		return iload_3
	case 0x1e:
		return lload_0
	case 0x1f:
		return lload_1
	case 0x20:
		return lload_2
	case 0x21:
		return lload_3
	case 0x22:
		return fload_0
	case 0x23:
		return fload_1
	case 0x24:
		return fload_2
	case 0x25:
		return fload_3
	case 0x26:
		return dload_0
	case 0x27:
		return dload_1
	case 0x28:
		return dload_2
	case 0x29:
		return dload_3
	case 0x2a:
		return aload_0
	case 0x2b:
		return aload_1
	case 0x2c:
		return aload_2
	case 0x2d:
		return aload_3
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
