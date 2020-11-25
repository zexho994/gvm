package classfile

// String结构的 tag 为 ConstantString
// stringIndex 是常量池中的有效索引，且索引处的类型必需是 CONSTANT_Utf8_info 类型
type ConstantStringInfo struct {
	tag         uint8
	cp          ConstantPool
	stringIndex uint16
}

/*
读取字符串的常量池索引
*/
func (constantStringInfo *ConstantStringInfo) readInfo(reader *ClassReader) {
	constantStringInfo.stringIndex = reader.readUint16()
}

/*
输出常量池中字符串的值
*/
func (constantStringInfo *ConstantStringInfo) String() string {
	return constantStringInfo.cp.getUtf8(constantStringInfo.stringIndex)
}
