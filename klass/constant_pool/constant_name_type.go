package constant_pool

import "github.com/zouzhihao-994/gvm/loader"

// 表示字段或者方法
type ConstantNameAndTypeInfo struct {
	Tag uint8
	// 常量池索引，对应 ConstantUtf8Info 结构
	// 表示方法｜字段名称，或者表示方法 <init>
	NameIndex uint16
	// 常量池索引，对应 ConstantUtf8Info 结构
	// 表示一个字段描述符或者方法描述符
	DescriptorIndex uint16
	cp              ConstantPool
}

func (constantNameAndTypeInfo *ConstantNameAndTypeInfo) ReadInfo(reader *loader.ClassReader) {
	constantNameAndTypeInfo.NameIndex = reader.ReadUint16()
	constantNameAndTypeInfo.DescriptorIndex = reader.ReadUint16()
}
