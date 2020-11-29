package jclass

import (
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
)

type ConstantPool []constant_pool.ConstantType

const (
	CONSTANT_Utf8               = 0x01 // utf8 字符串
	CONSTANT_Interger           = 0x03 // 整形常量，4 bytes
	CONSTANT_Float              = 0x04 // 浮点常量，4 bytes
	CONSTANT_Long               = 0x05 // 长整形常量，8 bytes
	CONSTANT_DOUBLE             = 0x06 // 双精度浮点常量，8 bytes
	CONSTANT_CLASS              = 0x07 // 类常量
	CONSTANT_String             = 0x08 // 字符串常量
	CONSTANT_Fieldref           = 0x09 // 字段的符号引用
	CONSTANT_InterfaceMethodref = 0x0b // 类方法的符号引用
	CONSANT_Methodref           = 0x0a // 接口方法的符号引用
	CONSTANT_NameAndType        = 0x0c // 接口方法的符号引用
	CONSTANT_MethodHandle       = 0x0f
	CONSTANT_MethodType         = 0x10
	CONSAANT_InvokeDynamic      = 0x12
)

/*
根据tag类型创建匹配的结构
*/
func newConstantInfo(tag uint8, cp ConstantPool) constant_pool.ConstantType {
	switch tag {
	case CONSTANT_Interger:
		return &constant_pool.ConstantIntegerInfo{Tag: tag}
	case CONSTANT_Float:
		return &constant_pool.ConstantFloatInfo{Tag: tag}
	case CONSTANT_Long:
		return &constant_pool.ConstantLongInfo{Tag: tag}
	case CONSTANT_DOUBLE:
		return &constant_pool.ConstantDoubleInfo{Tag: tag}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{tag: tag}
	case CONSTANT_String:
		return &ConstantStringInfo{tag: tag, cp: cp}
	case CONSTANT_CLASS:
		return &ConstantClassInfo{tag: tag, cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{tag: tag, cp: cp}}
	case CONSANT_Methodref:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{tag: tag, cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{tag: tag, cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{tag: tag}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{tag: tag}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{tag: tag}
	case CONSAANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{tag: tag}
	default:
		panic("java.lang.ClassFormatError: constant_pool pool tag!")
	}
}
