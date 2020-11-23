package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (constantNameAndTypeInfo *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	constantNameAndTypeInfo.nameIndex = reader.readUint16()
	constantNameAndTypeInfo.descriptorIndex = reader.readUint16()
}
