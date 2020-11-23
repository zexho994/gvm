package classfile

type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

type ConstantFieldRefInfo struct{ ConstantMemberRefInfo }

type ConstantMethodRefInfo struct{ ConstantMemberRefInfo }

type ConstantInterfaceMethodRefInfo struct{ ConstantMemberRefInfo }

/*
读取数据
*/
func (ConstantMemberRefInfo *ConstantMemberRefInfo) readInfo(reader *ClassReader) {

	ConstantMemberRefInfo.classIndex = reader.readUint16()

	ConstantMemberRefInfo.nameAndTypeIndex = reader.readUint16()
}

/*
获取类名
*/
func (ConstantMemberRefInfo *ConstantMemberRefInfo) ClassName() string {

	return ConstantMemberRefInfo.cp.getClassName(ConstantMemberRefInfo.classIndex)
}

/*
获取描述符
*/
func (ConstantMemberRefInfo *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {

	return ConstantMemberRefInfo.cp.getNameAndType(ConstantMemberRefInfo.nameAndTypeIndex)
}
