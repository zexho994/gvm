package instructions

import (
	"../instructions/base"
	"../instructions/comparisons"
	"../instructions/constants"
	"../instructions/control"
	"../instructions/loads"
	"../instructions/math"
	"../instructions/references"
	"../instructions/stack"
	"../instructions/stores"
	"fmt"
)

var (
	nop             = &constants.NOP{}
	aconst_null     = &constants.ACONST_NULL{}
	iconst_m1       = &constants.ICONST_M1{}
	iconst_0        = &constants.ICONST_0{}
	iconst_1        = &constants.ICONST_1{}
	iconst_2        = &constants.ICONST_2{}
	iconst_3        = &constants.ICONST_3{}
	iconst_4        = &constants.ICONST_4{}
	iconst_5        = &constants.ICONST_5{}
	lconst_0        = &constants.LCONST_0{}
	lconst_1        = &constants.LCONST_1{}
	fconst_0        = &constants.FCONST_0{}
	fconst_1        = &constants.FCONST_1{}
	fconst_2        = &constants.FCONST_2{}
	dconst_0        = &constants.DCONST_0{}
	dconst_1        = &constants.DCONST_1{}
	iload_0         = &loads.ILOAD_0{}
	iload_1         = &loads.ILOAD_1{}
	iload_2         = &loads.ILOAD_2{}
	iload_3         = &loads.ILOAD_3{}
	lload_0         = &loads.LLOAD_0{}
	lload_1         = &loads.LLOAD_1{}
	lload_2         = &loads.LLOAD_2{}
	lload_3         = &loads.LLOAD_3{}
	fload_0         = &loads.FLOAD_0{}
	fload_1         = &loads.FLOAD_1{}
	fload_2         = &loads.FLOAD_2{}
	fload_3         = &loads.FLOAD_3{}
	dload_0         = &loads.DLOAD_0{}
	dload_1         = &loads.DLOAD_1{}
	dload_2         = &loads.DLOAD_2{}
	dload_3         = &loads.DLOAD_3{}
	aload_0         = &loads.ALOAD_0{}
	aload_1         = &loads.ALOAD_1{}
	aload_2         = &loads.ALOAD_2{}
	aload_3         = &loads.ALOAD_3{}
	istore_1        = &stores.ISTORE_1{}
	istore_2        = &stores.ISTORE_2{}
	iadd            = &math.IADD{}
	iinc            = &math.IINC{}
	goto_           = &control.GOTO{}
	if_icmpgt       = &comparisons.IF_ICMPGT{}
	bipush          = &constants.BIPUSH{}
	ldc             = &constants.LDC{}
	new             = &references.NEW{}
	invokespecial   = &references.INVOKE_SPECIAL{}
	invokevirtual   = &references.INVOKE_VIRTUAL{}
	invokeinterface = &references.INVOKE_INTERFACE{}
	invokestatic    = &references.INVOKE_STATIC{}
	putstatic       = &references.PUT_STATIC{}

	ireturn = &control.IRETURN{}
	lreturn = &control.LRETURN{}
	freturn = &control.FRETURN{}
	dreturn = &control.DRETURN{}
	areturn = &control.ARETURN{}
	_return = &control.RETURN{}

	getstatic  = &references.GET_STATIC{}
	putfield   = &references.PUT_FIELD{}
	getfield   = &references.GET_FIELD{}
	checkcast  = &references.CHECK_CAST{}
	instanceof = &references.INSTANCE_OF{}
	dup        = &stack.Dup{}
	astore     = &stores.ASTORE{}
	astore_1   = &stores.ASTORE_1{}
	astore_2   = &stores.ASTORE_2{}
	astore_3   = &stores.ASTORE_3{}
	ifeq       = &comparisons.IFEQ{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	case 0x02:
		return iconst_m1
	case 0x03:
		return iconst_0
	case 0x04:
		return iconst_1
	case 0x05:
		return iconst_2
	case 0x06:
		return iconst_3
	case 0x07:
		return iconst_4
	case 0x08:
		return iconst_5
	case 0x09:
		return lconst_0
	case 0x0a:
		return lconst_1
	case 0x0b:
		return fconst_0
	case 0x0c:
		return fconst_1
	case 0x0d:
		return fconst_2
	case 0x0e:
		return dconst_0
	case 0x0f:
		return dconst_1
	case 0x10:
		return bipush
	case 0x12:
		return ldc
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
	case 0x3c:
		return istore_1
	case 0x3d:
		return istore_2
	case 0x60:
		return iadd
	case 0x84:
		return iinc
	case 0xa3:
		return if_icmpgt
	case 0xa7:
		return goto_
	case 0xbb:
		return new
	case 0x59:
		return dup

	case 0xac:
		return ireturn
	case 0xad:
		return lreturn
	case 0xae:
		return freturn
	case 0xaf:
		return dreturn
	case 0xb0:
		return areturn
	case 0xb1:
		return _return
	case 0xb2:
		return getstatic
	case 0xb3:
		return putstatic
	case 0xb4:
		return getfield
	case 0xb5:
		return putfield
	case 0xb6:
		return invokespecial
	case 0xb7:
		return invokespecial
	case 0xb8:
		return invokestatic
	case 0xb9:
		return invokeinterface

	case 0xc1:
		return instanceof
	case 0xc0:
		return checkcast
	case 0x4d:
		return astore_2
	case 0x4e:
		return astore_3
	case 0x99:
		return ifeq

	default:
		panic(fmt.Errorf("Unsupported opcode : 0x%x!", opcode))

	}
}
