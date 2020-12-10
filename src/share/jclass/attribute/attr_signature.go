package attribute

import "github.com/zouzhihao-994/gvm/src/share/classfile"

// 可选的变长属性，位于ClassFile、field_info、method_info的属性表中
// 包含该字段表示的类，接口，构造器方法或者字段包含类型变量或者参数化类型
// signature会为它记录泛型签名信息
type Attr_Signature struct {
	nameIdx uint16
	name    string
	attrLen uint32
	// 常量池中 UTF8 类型的索引，表示类签名、方法类型签名或字段类型签名
	signatureIdx uint16
}

func (attr Attr_Signature) parse(reader *classfile.ClassReader) {
	attr.signatureIdx = reader.ReadUint16()
}

func (attr Attr_Signature) Name() string {
	return attr.name
}
