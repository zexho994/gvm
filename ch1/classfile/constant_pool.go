package classfile

/*
常量池
数组类型
*/
type ConstantPool []ConstantInfo

/*
常量类型
*/
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantPool(reader *ClassReader) ConstantPool{
	// 常量池
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1 ; i < cpCount;i++{

	}
}

/*
获取常量信息
*/
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	// tag是第一个标记
	tag := reader.readUint8()
	// 根据tag创建一个常量对象
	c := newConstantInfo(tag, cp)
	// 读取常量信息
	c.readInfo(reader)

	return c
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


