package constant_pool

import (
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/loader"
)

type ConstantPool []ConstantType

const (
	ConstantUtf8               = 0x01 // utf8 字符串
	ConstantInterger           = 0x03 // 整形常量，4 bytes
	ConstantFloat              = 0x04 // 浮点常量，4 bytes
	ConstantLong               = 0x05 // 长整形常量，8 bytes
	ConstantDouble             = 0x06 // 双精度浮点常量，8 bytes
	ConstantClass              = 0x07 // 类常量
	ConstantString             = 0x08 // 字符串常量
	ConstantFieldref           = 0x09 // 字段的符号引用
	ConstantInterfacemethodref = 0x0b // 类方法的符号引用
	ConsantMethodref           = 0x0a // 接口方法的符号引用
	ConstantNameandtype        = 0x0c // 接口方法的符号引用
	ConstantMethodhandle       = 0x0f
	ConstantMethodtype         = 0x10 // 表示方法类型
	ConsaantInvokedynamic      = 0x12
)

// NewConstantInfo 根据tag类型创建匹配的结构
func (pool *ConstantPool) NewConstantInfo(tag uint8) ConstantType {
	switch tag {
	case ConstantInterger:
		return &ConstantIntegerInfo{Tag: tag}
	case ConstantFloat:
		return &ConstantFloatInfo{Tag: tag}
	case ConstantLong:
		return &ConstantLongInfo{Tag: tag}
	case ConstantDouble:
		return &ConstantDoubleInfo{Tag: tag}
	case ConstantUtf8:
		return &ConstantUtf8Info{Tag: tag}
	case ConstantString:
		return &ConstantStringInfo{Tag: tag, ConstantPool: pool}
	case ConstantClass:
		return &ConstantClassInfo{Tag: tag, ConstantPool: pool}
	case ConstantFieldref:
		return &ConstantFieldInfo{Tag: tag, ConstantPool: pool}
	case ConsantMethodref:
		return &ConstantMethodInfo{Tag: tag, ConstantPool: pool}
	case ConstantInterfacemethodref:
		return &ConstantInterfaceMethodInfo{Tag: tag, ConstantPool: pool}
	case ConstantNameandtype:
		return &ConstantNameAndTypeInfo{Tag: tag, ConstantPool: pool}
	case ConstantMethodtype:
		return &ConstantMethodTypeInfo{Tag: tag, ConstantPool: pool}
	case ConstantMethodhandle:
		return &ConstantMethodHandleInfo{Tag: tag, ConstantPool: pool}
	case ConsaantInvokedynamic:
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
	utf8Info := pool.GetConstantInfo(idx).(*ConstantUtf8Info)
	return utf8Info.Str
}

func (pool ConstantPool) GetClassName(index uint16) string {
	classInfo := pool.GetConstantInfo(index).(*ConstantClassInfo)
	return pool.GetUtf8(classInfo.NameIdx)
}

func (pool ConstantPool) GetNameAndType(index uint16) (string, string) {
	ntInfo := pool.GetConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := pool.GetUtf8(ntInfo.NameIndex)
	desc := pool.GetUtf8(ntInfo.DescriptorIndex)
	return name, desc
}

// ReadConstantPool 读取常量池数据
// 解析常量池分为两步：分配内存 -> 解析
func ReadConstantPool(cpCount uint16, reader *loader.ClassReader) ConstantPool {
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
func readConstantInfo(reader *loader.ClassReader, cp ConstantPool) ConstantType {
	// tag是第一个标记
	tag := reader.ReadUint8()
	// 根据tag创建一个常量对象
	c := cp.NewConstantInfo(tag)
	// 调用常量的read()方法
	c.ReadInfo(reader)
	return c
}
