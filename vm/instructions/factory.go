package instructions

import (
	"../instructions/base"
	"../instructions/comparisons"
	"../instructions/constants"
	"../instructions/loads"
	"../instructions/stores"
	"fmt"
)

var (
	nop         = &constants.NOP{}
	aconst_null = &constants.ACONST_NULL{}
	iconst_m1   = &constants.ICONST_M1{}
	iconst_0    = &constants.ICONST_0{}
	iconst_1    = &constants.ICONST_1{}
	iconst_2    = &constants.ICONST_2{}
	iconst_3    = &constants.ICONST_3{}
	iconst_4    = &constants.ICONST_4{}
	iconst_5    = &constants.ICONST_5{}
	lconst_0    = &constants.LCONST_0{}
	lconst_1    = &constants.LCONST_1{}
	fconst_0    = &constants.FCONST_0{}
	fconst_1    = &constants.FCONST_1{}
	fconst_2    = &constants.FCONST_2{}
	dconst_0    = &constants.DCONST_0{}
	dconst_1    = &constants.DCONST_1{}
	iload_0     = &loads.ILOAD_0{}
	iload_1     = &loads.ILOAD_1{}
	iload_2     = &loads.ILOAD_2{}
	iload_3     = &loads.ILOAD_3{}
	lload_0     = &loads.LLOAD_0{}
	lload_1     = &loads.LLOAD_1{}
	lload_2     = &loads.LLOAD_2{}
	lload_3     = &loads.LLOAD_3{}
	fload_0     = &loads.FLOAD_0{}
	fload_1     = &loads.FLOAD_1{}
	fload_2     = &loads.FLOAD_2{}
	fload_3     = &loads.FLOAD_3{}
	dload_0     = &loads.DLOAD_0{}
	dload_1     = &loads.DLOAD_1{}
	dload_2     = &loads.DLOAD_2{}
	dload_3     = &loads.DLOAD_3{}
	aload_0     = &loads.ALOAD_0{}
	aload_1     = &loads.ALOAD_1{}
	aload_2     = &loads.ALOAD_2{}
	aload_3     = &loads.ALOAD_3{}
	istore_1    = &stores.ISTORE_1{}
	istore_2    = &stores.ISTORE_2{}
	//iaload        =
	//laload        =
	//faload        =
	//daload        =
	//aaload        =
	//baload        =
	//caload        =
	//saload        =
	//istore_0      =
	//istore_1      =
	//istore_2      =
	//istore_3      =
	//lstore_0      =
	//lstore_1      =
	//lstore_2      =
	//lstore_3      =
	//fstore_0      =
	//fstore_1      =
	//fstore_2      =
	//fstore_3      =
	//dstore_0      =
	//dstore_1      =
	//dstore_2      =
	//dstore_3      =
	//astore_0      =
	//astore_1      =
	//astore_2      =
	//astore_3      =
	//iastore       =
	//lastore       =
	//fastore       =
	//dastore       =
	//aastore       =
	//bastore       =
	//castore       =
	//sastore       =
	//pop           =
	//pop2          =
	//dup           =
	//dup_x1        =
	//dup_x2        =
	//dup2          =
	//dup2_x1       =
	//dup2_x2       =
	//swap          =
	//iadd          =
	//ladd          =
	//fadd          =
	//dadd          =
	//isub          =
	//lsub          =
	//fsub          =
	//dsub          =
	//imul          =
	//lmul          =
	//fmul          =
	//dmul          =
	//idiv          =
	//ldiv          =
	//fdiv          =
	//ddiv          =
	//irem          =
	//lrem          =
	//frem          =
	//drem          =
	//ineg          =
	//lneg          =
	//fneg          =
	//dneg          =
	//ishl          =
	//lshl          =
	//ishr          =
	//lshr          =
	//iushr         =
	//lushr         =
	//iand          =
	//land          =
	//ior           =
	//lor           =
	//ixor          =
	//lxor          =
	//i2l           =
	//i2f           =
	//i2d           =
	//l2i           =
	//l2f           =
	//l2d           =
	//f2i           =
	//f2l           =
	//f2d           =
	//d2i           =
	//d2l           =
	//d2f           =
	//i2b           =
	//i2c           =
	//i2s           =
	//lcmp          =
	//fcmpl         =
	//fcmpg         =
	//dcmpl         =
	//dcmpg         =
	//ireturn       =
	//lreturn       =
	//freturn       =
	//dreturn       =
	//areturn       =
	//_return       =
	//arraylength   =
	//athrow        =
	//monitorenter  =
	//monitorexit   =
	//invoke_native =
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
		return &constants.BIPUSH{}
	//case 0x11:
	//	return &SIPush{}
	//case 0x12:
	//	return &LDC{}
	//case 0x13:
	//	return &LDC_W{}
	//case 0x14:
	//	return &LDC2_W{}
	//case 0x15:
	//	return NewLoad(false)
	//case 0x16:
	//	return NewLoad(true)
	//case 0x17:
	//	return NewLoad(false)
	//case 0x18:
	//	return NewLoad(true)
	//case 0x19:
	//	return NewLoad(false)
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
	case 0xa3:
		return &comparisons.IF_ICMPGT{}

	default:
		panic(fmt.Errorf("Unsupported opcode : 0x%x!", opcode))

	}
}
