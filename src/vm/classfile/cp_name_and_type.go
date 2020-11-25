package classfile

type ConstantNameAndTypeInfo struct {
	tag             uint8
	nameIndex       uint16
	descriptorIndex uint16
}

func (constantNameAndTypeInfo *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	constantNameAndTypeInfo.nameIndex = reader.readUint16()
	constantNameAndTypeInfo.descriptorIndex = reader.readUint16()
}
