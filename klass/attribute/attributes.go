package attribute

import (
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/loader"
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
	parse(reader *loader.ClassReader)
}

func ParseAttributes(attrCount uint16, reader *loader.ClassReader, cp constant_pool.ConstantPool) AttributesInfo {
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
		return &AttrCode{NameIdx: nameIdx, name: name, AttrLen: attrLen, cp: cp}
	case "ConstantValue":
		return &AttrConstantvalue{nameIdx: nameIdx, name: name, attrLen: attrLen, cp: cp}
	case "Exceptions":
		return &AttrExceptions{nameIdx: nameIdx, name: name, attrlen: attrLen, cp: cp}
	case "LineNumberTable":
		return &AttrLinenumbertable{nameIdx: nameIdx, name: name, cp: cp}
	case "SourceFile":
		return &AttrSourcefile{nameIdx: nameIdx, name: name, attrLen: attrLen}
	case "Signature":
		return &AttrSignature{nameIdx: nameIdx, name: name, attrLen: attrLen}
	case "StackMapTable":
		return &AttrStackmaptable{nameIdx: nameIdx, name: name, attrLen: attrLen}
	case "Deprecated":
		return &AttrDeprecated{nameIdx: nameIdx, name: name, attrLen: attrLen}
	case "RuntimeVisibleAnnotations":
		return &AttrRuntimevisibleannotation{nameIdx: nameIdx, name: name, attrLen: attrLen}
	//case "LocalVariableTable":
	//	return &LocalVariableTableAttribute{}
	case "Synthetic":
		return &AttrSynthetic{nameIdx: nameIdx, name: name, attrLen: attrLen}
	case "InnerClasses":
		return &AttrInnerClasses{nameIdx: nameIdx, name: name, attrLen: attrLen, cp: cp}
	case "BootstrapMethods":
		return &BootstrapmethodsAttribute{nameIdx: nameIdx, name: name, attrLen: attrLen}
	default:
		panic("attribute error")
	}
}

func (attrs AttributesInfo) AttrCode() (*AttrCode, error) {
	for idx := range attrs {
		if attrs[idx].Name() == "Code" {
			a := attrs[idx]
			code := a.(*AttrCode)
			return code, nil
		}
	}
	return nil, exception.GvmError{Msg: "method not exist the attribute of code"}
}
