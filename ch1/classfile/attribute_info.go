package classfile

type AttributeInfo interface{ readInfo(reader *ClassReader) }

/*
读取属性表
*/
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	// 读取2字节
	attributesCount := reader.readUint16()
	// 2字节大小的AttributeInfo数组
	attributes := make([]AttributeInfo, attributesCount)
	// 遍历数组
	for i := range attributes {
		//
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	// 读取2字节
	attrNameIndex := reader.readUint16()
	// 属性名:获取该所以对应常量池中的utf8数据
	attrName := cp.getUtf8(attrNameIndex)
	// 读取4自己诶
	attrLen := reader.readUint32()
	// 创建具体的属性实例
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	//
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
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
