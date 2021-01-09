package constant_pool

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	"github.com/zouzhihao-994/gvm/src/share/exception"
)

type ConstantPool []ConstantType

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
	CONSTANT_MethodType         = 0x10 // 表示方法类型
	CONSAANT_InvokeDynamic      = 0x12
)

/*
根据tag类型创建匹配的结构
*/
func (pool ConstantPool) NewConstantInfo(tag uint8) ConstantType {
	switch tag {
	case CONSTANT_Interger:
		return &ConstantIntegerInfo{Tag: tag}
	case CONSTANT_Float:
		return &ConstantFloatInfo{Tag: tag}
	case CONSTANT_Long:
		return &ConstantLongInfo{Tag: tag}
	case CONSTANT_DOUBLE:
		return &ConstantDoubleInfo{Tag: tag}
	case CONSTANT_Utf8:
		return &ConstantUtf8{Tag: tag}
	case CONSTANT_String:
		return &ConstantStringInfo{Tag: tag, Cp: pool}
	case CONSTANT_CLASS:
		return &ConstantClassInfo{Tag: tag, Cp: pool}
	case CONSTANT_Fieldref:
		return &ConstantFieldRefInfo{Tag: tag, Cp: pool}
	case CONSANT_Methodref:
		return &ConstantMethodInfo{Tag: tag, Cp: pool}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethod{Tag: tag, Cp: pool}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{Tag: tag, cp: pool}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{Tag: tag, Cp: pool}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{Tag: tag}
	case CONSAANT_InvokeDynamic:
		return &ConstantInvokeDynamic{Tag: tag}
	default:
		exception.GvmError{Msg: "java.lang.ClassFormatError: constant_pool pool tag!"}.Throw()
		return nil
	}
}

func (pool ConstantPool) GetConstantInfo(idx uint16) ConstantType {
	if info := pool[idx]; info != nil {
		return info
	}
	panic("[gvm] Invalid constant_pool index!")
}

func (pool ConstantPool) GetUtf8(idx uint16) string {
	utf8Info := pool.GetConstantInfo(idx).(*ConstantUtf8)
	return utf8Info.Str
}

func (pool ConstantPool) GetClassName(index uint16) string {
	classInfo := pool.GetConstantInfo(index).(*ConstantClassInfo)
	return pool.GetUtf8(classInfo.NameIdx)
}

func (pool ConstantPool) GetNameAndType(index uint16) (string, string) {
	ntInfo := pool.GetConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := pool.GetUtf8(ntInfo.NameIndex)
	_type := pool.GetUtf8(ntInfo.DescriptorIndex)
	return name, _type
}

// 读取常量池数据
// 解析常量池分为两步：分配内存 -> 解析
func ReadConstantPool(cpCount uint16, reader *classfile.ClassReader) ConstantPool {
	cp := make([]ConstantType, cpCount)
	for i := uint16(1); i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

/*
获取常量信息
1. 获取tag , 调用newConstantInfo()创建具体的常量
2. 调用readInfo()方法读取常量信息
*/
func readConstantInfo(reader *classfile.ClassReader, cp ConstantPool) ConstantType {
	// tag是第一个标记
	tag := reader.ReadUint8()
	// 根据tag创建一个常量对象
	c := cp.NewConstantInfo(tag)
	// 调用常量的read()方法
	c.ReadInfo(reader)
	return c
}
