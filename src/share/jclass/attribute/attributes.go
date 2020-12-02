package attribute

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
)

type AttributeInfo interface {
	parse(reader *classfile.ClassReader)
}

func ParseAttributes(attrCount uint16, reader *classfile.ClassReader, cp constant_pool.ConstantPool) []AttributeInfo {
	attributes := make([]AttributeInfo, attrCount)
	for i := range attributes {
		attrNameIdx := reader.ReadUint16()
		attrLen := reader.ReadUint32()
		attrInfo := newAttributeInfo(attrNameIdx, attrLen, cp)
		attrInfo.parse(reader)
		attributes[i] = attrInfo
	}
	return attributes
}

func newAttributeInfo(nameIdx uint16, attrLen uint32, cp constant_pool.ConstantPool) AttributeInfo {
	name := cp.GetUtf8(nameIdx)
	switch name {
	case "Code":
		return &Attr_Code{NameIdx: nameIdx, AttrLen: attrLen, cp: cp}
	case "ConstantValue":
		return &Attr_ConstantValue{nameIdx: nameIdx, name: name, attrLen: attrLen, cp: cp}
	case "Exceptions":
		return &Attr_Exceptions{nameIdx: nameIdx, name: name, attrlen: attrLen, cp: cp}
	case "LineNumberTable":
		return &Attr_LineNumberTable{nameIdx: nameIdx, name: name, cp: cp}
	case "SourceFile":
		return &Attr_SourceFile{nameIdx: nameIdx, name: name, attrLen: attrLen}
	case "Signature":
		return &Attr_Signature{nameIdx: nameIdx, name: name, attrLen: attrLen}
	case "StackMapTable":
	//case "Deprecated":
	//	return &DeprecatedAttribute{}
	//case "LocalVariableTable":
	//	return &LocalVariableTableAttribute{}
	//case "Synthetic":
	//	return &SyntheticAttribute{}

	default:
		panic("attribute error")
		//return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
