package attribute

import (
	"github.com/zouzhihao-994/gvm/classloader"
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/jclass/constant_pool"
)

type AttributesInfo []IAttributeInfo

func (attrs AttributesInfo) FindAttrInfo(name string) (IAttributeInfo, error) {
	for idx := range attrs {
		if attrs[idx].Name() == name {
			return attrs[idx], nil
		}
	}
	return nil, exception.GvmError{Msg: exception.AttributeNotFoundError}
}

type IAttributeInfo interface {
	Name() string
	parse(reader *classloader.ClassReader)
}

func ParseAttributes(attrCount uint16, reader *classloader.ClassReader, cp constant_pool.ConstantPool) AttributesInfo {
	attributes := make(AttributesInfo, attrCount)
	for i := range attributes {
		attrNameIdx := reader.ReadUint16()
		attrLen := reader.ReadUint32()
		attrInfo := newAttributeInfo(attrNameIdx, attrLen, cp)
		attrInfo.parse(reader)
		attributes[i] = attrInfo
	}
	return attributes
}

func newAttributeInfo(nameIdx uint16, attrLen uint32, cp constant_pool.ConstantPool) IAttributeInfo {
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
		return &BootstrapmethodsAttribute{nameIdx: nameIdx, name: name, attrLen: attrLen}
	default:
		panic("attribute error")
	}
}

func (attrs AttributesInfo) AttrCode() (*Attr_Code, error) {
	for idx := range attrs {
		if attrs[idx].Name() == "Code" {
			a := attrs[idx]
			code := a.(*Attr_Code)
			return code, nil
		}
	}
	return nil, exception.GvmError{Msg: "method not exist the attribute of code"}
}
