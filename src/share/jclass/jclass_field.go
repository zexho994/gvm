package jclass

/*
字段表
*/
type FieldInfo struct {
	// 常量池指针
	constantPool Constant
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
