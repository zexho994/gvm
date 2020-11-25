package classfile

/*
常量池
数组类型
*/
type ConstantPool []ConstantInfo

/*
常量信息
*/
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

// 读取常量池数据
// 解析常量池分为两步：分配内存 -> 解析
func readConstantPool(cpCount uint16, reader *ClassReader) ConstantPool {
	cp := make([]ConstantInfo, cpCount)
	for i := uint16(1); i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		// long or double need two bytes
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
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	// tag是第一个标记
	tag := reader.readUint8()
	// 根据tag创建一个常量对象
	c := newConstantInfo(tag, cp)
	// 调用常量的read()方法
	c.readInfo(reader)
	return c
}

/*
按照索引(index)查找常量
*/
func (constantPool ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := constantPool[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (constantPool ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := constantPool.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := constantPool.getUtf8(ntInfo.nameIndex)
	_type := constantPool.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (constantPool ConstantPool) getClassName(index uint16) string {
	classInfo := constantPool.getConstantInfo(index).(*ConstantClassInfo)
	return constantPool.getUtf8(classInfo.nameIndex)
}

/*
从常量池中查找UTF-8字符串
*/
func (constantPool ConstantPool) getUtf8(index uint16) string {
	utf8Info := constantPool.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

/*
根据tag类型创建匹配的结构
*/
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case ConstantInteger:
		return &ConstantIntegerInfo{tag: tag}
	case ConstantFloat:
		return &ConstantFloatInfo{tag: tag}
	case ConstantLong:
		return &ConstantLongInfo{tag: tag}
	case ConstantDouble:
		return &ConstantDoubleInfo{tag: tag}
	case ConstantUtf8:
		return &ConstantUtf8Info{tag: tag}
	case ConstantString:
		return &ConstantStringInfo{tag: tag, cp: cp}
	case ConstantClass:
		return &ConstantClassInfo{tag: tag, cp: cp}
	case ConstantFieldRef:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{tag: tag, cp: cp}}
	case ConstantMethodRef:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{tag: tag, cp: cp}}
	case ConstantInterfaceMethodRef:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{tag: tag, cp: cp}}
	case ConstantNameAndType:
		return &ConstantNameAndTypeInfo{tag: tag}
	case ConstantMethodType:
		return &ConstantMethodTypeInfo{tag: tag}
	case ConstantMethodHandle:
		return &ConstantMethodHandleInfo{tag: tag}
	case ConstantInvokeDynamic:
		return &ConstantInvokeDynamicInfo{tag: tag}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
