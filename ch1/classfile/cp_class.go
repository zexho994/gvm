package classfile

/*
常量池中类信息结构
*/
type ConstantClassInfo struct {
	// 指向常量池的指针
	cp ConstantPool
	// 下标
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {

	self.nameIndex = reader.readUint16()
}
func (self *ConstantClassInfo) Name() string {

	return self.cp.getUtf8(self.nameIndex)
}
