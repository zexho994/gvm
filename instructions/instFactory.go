package instructions

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/instructions/comparisons"
	constants "github.com/zouzhihao-994/gvm/instructions/constant"
	"github.com/zouzhihao-994/gvm/instructions/control"
	"github.com/zouzhihao-994/gvm/instructions/conversions"
	"github.com/zouzhihao-994/gvm/instructions/loads"
	"github.com/zouzhihao-994/gvm/instructions/math"
	"github.com/zouzhihao-994/gvm/instructions/references"
	"github.com/zouzhihao-994/gvm/instructions/reserved"
	"github.com/zouzhihao-994/gvm/instructions/stack"
	"github.com/zouzhihao-994/gvm/instructions/stores"
)

var (
	nop        = &base.NOP{}
	aconstNull = &constants.AconstNull{}
	iconstMl   = &constants.IconstM1{}
	iconst0    = &constants.Iconst0{}
	iconst1    = &constants.Iconst1{}
	iconst2    = &constants.Iconst2{}
	iconst3    = &constants.Iconst3{}
	iconst4    = &constants.Iconst4{}
	iconst5    = &constants.Iconst5{}
	lconst0    = &constants.Lconst0{}
	lconst1    = &constants.Lconst1{}
	fconst0    = &constants.Fconst0{}
	fconst1    = &constants.Fconst1{}
	fconst2    = &constants.Fconst2{}
	dconst0    = &constants.Dconst0{}
	dconst1    = &constants.Dconst1{}

	ldc    = &constants.LDC{}
	ldc2W  = &constants.Ldc2W{}
	bipush = &constants.BIPUSH{}
	sipush = &constants.SIPUSH{}

	iload  = &loads.ILOAD{}
	lload  = &loads.LLOAD{}
	fload  = &loads.FLOAD{}
	dload  = &loads.DLOAD{}
	aload  = &loads.ALOAD{}
	iload0 = &loads.Iload0{}
	iload1 = &loads.Iload1{}
	iload2 = &loads.Iload2{}
	iload3 = &loads.Iload3{}
	lload0 = &loads.Lload0{}
	lload1 = &loads.Lload1{}
	lload2 = &loads.Lload2{}
	lload3 = &loads.Lload3{}
	fload0 = &loads.Fload0{}
	fload1 = &loads.Fload1{}
	fload2 = &loads.Fload2{}
	fload3 = &loads.Fload3{}
	dload0 = &loads.Dload0{}
	dload1 = &loads.Dload1{}
	dload2 = &loads.Dload2{}
	dload3 = &loads.Dload3{}
	aload0 = &loads.Aload0{}
	aload1 = &loads.Aload1{}
	aload2 = &loads.Aload2{}
	aload3 = &loads.Aload3{}

	istore  = &stores.ISTORE{}
	istore0 = &stores.ISTORE_0{}
	istore1 = &stores.ISTORE_1{}
	istore2 = &stores.ISTORE_2{}
	istore3 = &stores.ISTORE_3{}
	astore  = &stores.ASTORE{}
	astore0 = &stores.Astore0{}
	astore1 = &stores.Astore1{}
	astore2 = &stores.Astore2{}
	astore3 = &stores.Astore3{}
	fstore  = &stores.FSTORE{}
	fstore0 = &stores.FSTORE_0{}
	fstore1 = &stores.FSTORE_1{}
	fstore2 = &stores.FSTORE_2{}
	fstore3 = &stores.FSTORE_3{}
	dstore  = &stores.DSTORE{}
	dstore0 = &stores.Dstore0{}
	dstore1 = &stores.Dstore1{}
	dstore2 = &stores.Dstore2{}
	dstore3 = &stores.Dstore3{}
	lstore  = &stores.LSTORE{}
	lstore0 = &stores.LSTORE_0{}
	lstore1 = &stores.LSTORE_1{}
	lstore2 = &stores.LSTORE_2{}
	lstore3 = &stores.LSTORE_3{}
	iastore = &stores.IASTORE{}
	sastore = &stores.SASTORE{}

	iadd = &math.IADD{}
	ladd = &math.LADD{}
	dadd = &math.DADD{}
	fadd = &math.FADD{}

	isub = &math.ISUB{}
	lsub = &math.LSUB{}
	fsub = &math.FSUB{}
	dsub = &math.DSUB{}

	imul = &math.IMUL{}
	lmul = &math.LMUL{}
	fmul = &math.FMUL{}
	dmul = &math.DMUL{}

	idiv = &math.IDIV{}
	ldiv = &math.LDIV{}
	ddiv = &math.DDIV{}
	fdiv = &math.FDIV{}

	_return  = &control.RETURN{}
	_ireturn = &control.IRETURN{}
	_areturn = &control.ARETURN{}
	_dreturn = &control.DRETURN{}
	_lreturn = &control.LRETURN{}
	_freturn = &control.FRETURN{}

	pop  = &stack.POP{}
	pop2 = &stack.POP2{}

	dup    = &stack.Dup{}
	dupX1  = &stack.DupX1{}
	dupX2  = &stack.DupX2{}
	dup2   = &stack.Dup2{}
	dup2X1 = &stack.Dup2X1{}
	dup2X2 = &stack.Dup2X2{}

	i2f = &conversions.I2f{}
	f2i = &conversions.F2i{}

	iinc  = &math.IINC{}
	_goto = &control.GOTO{}

	monitorEnter = &references.MonitorEntry{}
	monitorExit  = &references.MonitorExit{}

	lcmp      = &comparisons.LCMP{}
	fcmpg     = &comparisons.FCMPG{}
	fcmpl     = &comparisons.FCMPL{}
	ifIcmpge  = &comparisons.IfIcmpge{}
	ifIcmple  = &comparisons.IfIcmple{}
	ifIcmpne  = &comparisons.IfAcmpne{}
	ifIcmpeq  = &comparisons.IfAcmpeq{}
	ifIcmplt  = &comparisons.IfIcmplt{}
	ifIcmpgt  = &comparisons.IfIcmpgt{}
	ifAcmpeq  = &comparisons.IfAcmpeq{}
	ifAcmpne  = &comparisons.IfAcmpne{}
	ifNull    = &comparisons.IfNull{}
	ifNonnull = &comparisons.IfNonnull{}
	ifge      = &comparisons.IfGe{}
	ifle      = &comparisons.IfLe{}
	ifne      = &comparisons.IfNe{}
	ifgt      = &comparisons.IfGt{}
	ifeq      = &comparisons.IfEq{}
	iflt      = &comparisons.IfLt{}

	getStatic       = &references.GetStatic{} // 178
	putStatic       = &references.PutStatic{}
	getField        = &references.GetField{}
	putField        = &references.PutField{}
	invokeStatic    = &references.InvokeStatic{} // 184
	invokeSpecial   = &references.InvokeSpecial{}
	invokeVirtual   = &references.InvokeVirtual{}
	invokeDynamic   = &references.InvokeDynamic{}
	invokeInterface = &references.InvokeInterface{}
	invokeNative    = &reserved.InvokeNative{}
	_new            = &references.NEW{} //187
	anewarray       = &references.AnewArray{}
	_newArray       = &references.NewArray{}
	arrayLength     = &references.ArrayLength{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconstNull
	case 0x02:
		return iconstMl
	case 0x03:
		return iconst0
	case 0x04:
		return iconst1
	case 0x05:
		return iconst2
	case 0x06:
		return iconst3
	case 0x07:
		return iconst4
	case 0x08:
		return iconst5
	case 0x09:
		return lconst0
	case 0x0a:
		return lconst1
	case 0x0b:
		return fconst0
	case 0x0c:
		return fconst1
	case 0x0d:
		return fconst2
	case 0x0e:
		return dconst0
	case 0x0f:
		return dconst1
	case 0x10:
		return bipush
	case 0x11:
		return sipush
	case 0x12:
		return ldc
	//case 0x13:
	//	return ldc_w
	case 0x14:
		return ldc2W
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
		return iload0
	case 0x1b:
		return iload1
	case 0x1c:
		return iload2
	case 0x1d:
		return iload3
	case 0x1e:
		return lload0
	case 0x1f:
		return lload1
	case 0x20:
		return lload2
	case 0x21:
		return lload3
	case 0x22:
		return fload0
	case 0x23:
		return fload1
	case 0x24:
		return fload2
	case 0x25:
		return fload3
	case 0x26:
		return dload0
	case 0x27:
		return dload1
	case 0x28:
		return dload2
	case 0x29:
		return dload3
	case 0x2a:
		return aload0
	case 0x2b:
		return aload1
	case 0x2c:
		return aload2
	case 0x2d:
		return aload3
	//case 0x2e:
	//	return iadload
	//case 0x2f:
	//	return laload
	//case 0x30:
	//	return faload
	//case 0x31:
	//	return daload
	//case 0x32:
	//	return aaload
	//case 0x33:
	//	return baload
	//case 0x34:
	//	return caload
	//case 0x35:
	//	return saload
	case 0x36:
		return istore
	case 0x37:
		return lstore
	case 0x38:
		return fstore
	case 0x39:
		return dstore
	case 0x3a:
		return astore
	case 0x3b:
		return istore0
	case 0x3c:
		return istore1
	case 0x3d:
		return istore2
	case 0x3e:
		return istore3
	case 0x3f:
		return lstore0
	case 0x40:
		return lstore1
	case 0x41:
		return lstore2
	case 0x42:
		return lstore3
	case 0x43:
		return fstore0
	case 0x44:
		return fstore1
	case 0x45:
		return fstore2
	case 0x46:
		return fstore3
	case 0x47:
		return dstore0
	case 0x48:
		return dstore1
	case 0x49:
		return dstore2
	case 0x4a:
		return dstore3
	case 0x4b:
		return astore0
	case 0x4c:
		return astore1
	case 0x4d:
		return astore2
	case 0x4e:
		return astore3
	case 0x4f:
		return iastore
	//case 0x50:
	//	return lastore
	//case 0x51:
	//	return fastore
	//case 0x52:
	//	return dastore
	//case 0x53:
	//	return aastore
	//case 0x54:
	//	return bastore
	case 0x55:
		return sastore
	case 0x57:
		return pop
	case 0x58:
		return pop2
	case 0x59:
		return dup
	case 0x5a:
		return dupX1
	case 0x5b:
		return dupX2
	case 0x5c:
		return dup2
	case 0x5d:
		return dup2X1
	case 0x5e:
		return dup2X2
	case 0x60:
		return iadd
	case 0x61:
		return ladd
	case 0x62:
		return fadd
	case 0x63:
		return dadd
	case 0x64:
		return isub
	case 0x65:
		return lsub
	case 0x66:
		return fsub
	case 0x67:
		return dsub
	case 0x68:
		return imul
	case 0x69:
		return lmul
	case 0x6a:
		return fmul
	case 0x6b:
		return dmul
	case 0x6c:
		return idiv
	case 0x6d:
		return ldiv
	case 0x6e:
		return fdiv
	case 0x6f:
		return ddiv
	//case 0x70:
	//	return irem
	//case 0x71:
	//	return lrem
	//case 0x72:
	//	return frem
	//case 0x73:
	//	return drem
	//case 0x74:
	//	return ineg
	//case 0x75:
	//	return lneg
	//case 0x76:
	//	return fneg
	//case 0x77:
	//	return dneg
	//case 0x78:
	//	return ishl
	//case 0x79:
	//	return lshl
	//case 0x7a:
	//	return ishr
	//case 0x7b:
	//	return lshr
	//case 0x7c:
	//	return iushr
	//case 0x7d:
	//	return lushr
	//case 0x7e:
	//	return iand
	//case 0x7f:
	//	return land
	//case 0x80:
	//	return ior
	//case 0x81:
	//	return lor
	//case 0x82:
	//	return ixor
	//case 0x83:
	//	return lxor
	case 0x84:
		return iinc
	//case 0x85:
	//	return i2l
	case 0x86:
		return i2f
	//case 0x87:
	//	return i2d
	//case 0x88:
	//	return l2i
	//case 0x89:
	//	return l2f
	//case 0x8a:
	//	return l2d
	case 0x8b:
		return f2i
	//case 0x8c:
	//	return fl2
	//case 0x8d:
	//	return f2d
	//case 0x8e:
	//	return d2i
	//case 0x8f:
	//	return d2l
	//case 0x90:
	//	return d2f
	//case 0x91:
	//	return i2b
	//case 0x92:
	//	return i2c
	//case 0x93:
	//	return i2s
	case 0x94:
		return lcmp
	case 0x95:
		return fcmpl
	case 0x96:
		return fcmpg
	//case 0x97:
	//	return dcmpl
	//case 0x98:
	//	return dcmpl
	case 0x99:
		return ifeq
	case 0x9a:
		return ifne
	case 0x9b:
		return iflt
	case 0x9c:
		return ifge
	case 0x9d:
		return ifgt
	case 0x9e:
		return ifle
	case 0x9f:
		return ifIcmpeq
	case 0xa0:
		return ifIcmpne
	case 0xa1:
		return ifIcmplt
	case 0xa2:
		return ifIcmpge
	case 0xa3:
		return ifIcmpgt
	case 0xa4:
		return ifIcmple
	case 0xa5:
		return ifAcmpeq
	case 0xa6:
		return ifAcmpne
	case 0xa7:
		return _goto
	//case 0xa8:
	//	return jsr
	//case 0xa9:
	//	return ret
	//case 0xaa:
	//	return tableswitch
	//case 0xab:
	//	return lookupswitch
	case 0xac:
		return _ireturn
	case 0xad:
		return _lreturn
	case 0xae:
		return _freturn
	case 0xaf:
		return _dreturn
	case 0xb0:
		return _areturn
	case 0xb1:
		return _return
	case 0xb2:
		return getStatic
	case 0xb3:
		return putStatic
	case 0xb4:
		return getField
	case 0xb5:
		return putField
	case 0xb6:
		return invokeVirtual
	case 0xb7:
		return invokeSpecial
	case 0xb8:
		return invokeStatic
	case 0xb9:
		return invokeInterface
	case 0xba:
		return invokeDynamic
	case 0xbb:
		return _new
	case 0xbc:
		return _newArray
	case 0xbd:
		return anewarray
	case 0xbe:
		return arrayLength
	//case 0xbf:
	//	return athrow
	//case 0xc0:
	//	return checkcast
	//case 0xc1:
	//	return instanceof
	case 0xc2:
		return monitorEnter
	case 0xc3:
		return monitorExit
	//case 0xc4:
	//	return wide
	//case 0xc5:
	//	return multianewarray
	case 0xc6:
		return ifNull
	case 0xc7:
		return ifNonnull
	//case 0xc8:
	//	return goto_w
	//case 0xc9:
	//	return jsr_w
	//case 0xca:
	//	return breakpoint
	case 0xfe:
		return invokeNative
	//case 0xff:
	//	return impdep2
	default:
		panic(fmt.Errorf("Unsupported opcode : 0x %x!", opcode))
	}
}
