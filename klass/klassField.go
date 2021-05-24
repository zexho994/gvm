package klass

import (
	"github.com/zouzhihao-994/gvm/klass/attribute"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/loader"
)

type Fields []FieldInfo

// FieldInfo /*
type FieldInfo struct {
	// 常量池指针
	*constant_pool.ConstantPool
	// 访问类型
	AccessFlags uint16
	// 字段名索引 -> 常量池
	NameIndex uint16
	// 描述符索引 -> 常量池
	DescriptorIndex uint16
	// 属性表
	AttributesCount uint16
	Attributes      attribute.AttributesInfo
}

// 解析字段表
// 解析可以分为两部分：1 基本结构的解析 2 属性表的解析
func parseFields(count uint16, reader *loader.ClassReader, cp *constant_pool.ConstantPool) Fields {
	fields := make([]FieldInfo, count)
	for i := range fields {
		field := FieldInfo{}
		// 解析base
		field.AccessFlags = reader.ReadUint16()
		field.NameIndex = reader.ReadUint16()
		field.DescriptorIndex = reader.ReadUint16()
		field.ConstantPool = cp
		// 解析属性表
		field.AttributesCount = reader.ReadUint16()
		field.Attributes = attribute.ParseAttributes(field.AttributesCount, reader, cp)
		fields[i] = field
	}
	return fields
}

func (field FieldInfo) Descriptor() string {
	return field.ConstantPool.GetUtf8(field.DescriptorIndex)
}

func (field FieldInfo) Name() string {
	return field.ConstantPool.GetUtf8(field.NameIndex)
}
