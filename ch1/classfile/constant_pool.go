package classfile

/*
常量池
数组类型
*/
type ConstantPool []ConstantInfo

/*
读取常量池信息
*/
func readConstantPool(reader *ClassReader) ConstantPool {
	// 读取2字节
	cpCount := int(reader.readUint16())
	// 切片
	cp := make([]ConstantInfo, cpCount)
	// 遍历常量项
	for i := 1; i < cpCount; i++ { // 注意索引从 1开始
		// 存
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		// 对于是constanlonginfo和constantDoubleinfo两者,长度还要加一
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ // 占两个位置
		}
	}
	return cp
}

/*

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

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
