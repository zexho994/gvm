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

/*
读取常量池数据
解析常量池分为两步：分配内存 -> 解析
*/
func readConstantPool(reader *ClassReader) ConstantPool {
	//fmt.Println("[gvm][readConstantPool] read constanpool ...")
	// get constantpool count
	cpCount := int(reader.readUint16())
	// 分配内存
	cp := make([]ConstantInfo, cpCount)

	// traverse cp
	for i := 1; i < cpCount; i++ {
		// read class file
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
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

/*
从常量池中查找UTF-8字符串
*/
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

/*
构造函数
根据tag创建常量信息
*/
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
