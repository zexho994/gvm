package jclass

type ConstantPool struct {
	tag  uint8
	data []uint8
}

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
