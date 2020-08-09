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

	/*
	   读取常量池信息
	*/
	func
	readConstantPool(reader * ClassReader)
	ConstantPool{
		// 读取2字节
		cpCount, := int(reader.readUint16())
		// 切片
		cp := make([]ConstantInfo, cpCount)
		// 遍历常量项
		for i := 1; i < cpCount; i++{ // 注意索引从 1开始
		// 存
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type){
		// 对于是constanlonginfo和constantDoubleinfo两者,长度还要加一
	case *ConstantLongInfo, *ConstantDoubleInfo:
		i++ // 占两个位置
	}
	}
		return cp
	}

	/*

	 */
	func(self ConstantPool) getConstantInfo(index
	uint16) ConstantInfo{
		if cpInfo := self[index]; cpInfo != nil{
		return cpInfo
	}
		panic("Invalid constant pool index!")
	}

	func(self ConstantPool) getNameAndType(index
	uint16) (string, string) {
		ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
		name := self.getUtf8(ntInfo.nameIndex)
		_type := self.getUtf8(ntInfo.descriptorIndex)
		return name, _type
	}

	func(self ConstantPool) getClassName(index
	uint16) string{
		classInfo, := self.getConstantInfo(index).(*ConstantClassInfo)
		return self.getUtf8(classInfo.nameIndex)
	}

	func(self ConstantPool) getUtf8(index
	uint16) string{
		utf8Info, := self.getConstantInfo(index).(*ConstantUtf8Info)
		return utf8Info.str
	}
