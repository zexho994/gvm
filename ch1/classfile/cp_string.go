package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

/*
读取字符串的常量池索引
*/
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

/*
输出常量池中字符串的值
*/
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
