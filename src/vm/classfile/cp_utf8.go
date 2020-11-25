package classfile

/*
符号引用的指针都会指向常量池中的索引
索引对应的数据用本utf8结构显示
*/
type ConstantUtf8Info struct {
	tag uint8
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

/*
完整的在源码中补充
*/
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
