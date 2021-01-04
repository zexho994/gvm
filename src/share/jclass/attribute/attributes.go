package attribute

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
)

type AttributeInfos []AttributeInfo

type AttributeInfo interface {
	Name() string
	parse(reader *classfile.ClassReader)
}

func ParseAttributes(attrCount uint16, reader *classfile.ClassReader, cp constant_pool.ConstantPool) AttributeInfos {
	attributes := make(AttributeInfos, attrCount)
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
		return &Attr_Code{NameIdx: nameIdx, name: name, AttrLen: attrLen, cp: cp}
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
		return &Attr_StackMapTable{nameIdx: nameIdx, name: name, attrLen: attrLen}
	case "Deprecated":
		return &Attr_Deprecated{nameIdx: nameIdx, name: name, attrLen: attrLen}
	case "RuntimeVisibleAnnotations":
		return &Attr_RuntimeVisibleAnnotation{nameIdx: nameIdx, name: name, attrLen: attrLen}
	//case "LocalVariableTable":
	//	return &LocalVariableTableAttribute{}
	case "Synthetic":
		return &Attr_Synthetic{nameIdx: nameIdx, name: name, attrLen: attrLen}
	case "InnerClasses":
		return &Attr_InnerClasses{nameIdx: nameIdx, name: name, attrLen: attrLen, cp: cp}
	case "BootstrapMethods":
		return &BootstrapmethodsAttribute{}
	default:
		panic("attribute error")
	}
}

func (attrs AttributeInfos) AttrCode() (*Attr_Code, error) {
	for idx := range attrs {
		if attrs[idx].Name() == "Code" {
			a := attrs[idx]
			code := a.(*Attr_Code)
			return code, nil
		}
	}
	return nil, exception.GvmError{Msg: "method not exist the attribute of code"}
}
