package jclass

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
)

type Fields []FieldInfo

/*
字段表
*/
type FieldInfo struct {
	// 常量池指针
	constantPool constant_pool.ConstantPool
	// 访问类型
	accessFlags uint16
	// 字段名索引 -> 常量池
	nameIndex uint16
	// 描述符索引 -> 常量池
	descriptorIndex uint16
	// 属性表
	attributesCount uint16
	attributes      []AttributeInfo
}

// 解析字段表
// 解析可以分为两部分：1 基本结构的解析 2 属性表的解析
func parseFields(count uint16, reader classfile.ClassReader) Fields {
	fields := make([]FieldInfo, count)
	for idx := range fields {
		field := FieldInfo{}
		// 解析base
		field.accessFlags = reader.ReadUint16()
		field.nameIndex = reader.ReadUint16()
		field.descriptorIndex = reader.ReadUint16()
		// 解析属性表

	}
}

func parseBaseInfo() {

}
