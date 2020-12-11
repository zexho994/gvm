package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

/*
读取属性表
*/
func readAttributes(attributesCount uint16, reader *ClassReader, cp ConstantPool) []AttributeInfo {
	// 2字节大小的AttributeInfo数组
	attributes := make([]AttributeInfo, attributesCount)

	// 遍历数组
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}

	return attributes
}

//
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	// 名称索引
	attrNameIndex := reader.readUint16()

	// 名称
	attrName := cp.getUtf8(attrNameIndex)

	// 属性表长度
	attrLen := reader.readUint32()

	// 属性表信息
	attrInfo := newAttributeInfo(attrName, attrLen, cp)

	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "AttrCode":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
