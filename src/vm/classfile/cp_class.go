package classfile

/*
常量池中类信息结构
*/
type ConstantClassInfo struct {
	tag uint8
	// 指向常量池的指针
	cp ConstantPool
	// 下标
	nameIndex uint16
}

func (ConstantClassInfo *ConstantClassInfo) readInfo(reader *ClassReader) {

	ConstantClassInfo.nameIndex = reader.readUint16()
}
func (ConstantClassInfo *ConstantClassInfo) Name() string {

	return ConstantClassInfo.cp.getUtf8(ConstantClassInfo.nameIndex)
}
