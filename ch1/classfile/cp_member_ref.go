package classfile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }
type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }

/*
读取数据
*/
func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {

	self.classIndex = reader.readUint16()

	self.nameAndTypeIndex = reader.readUint16()
}

/*
获取类名
*/
func (self *ConstantMemberrefInfo) ClassName() string {

	return self.cp.getClassName(self.classIndex)
}

/*
获取描述符
*/
func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {

	return self.cp.getNameAndType(self.nameAndTypeIndex)
}
