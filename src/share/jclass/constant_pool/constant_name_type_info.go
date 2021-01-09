package constant_pool

import "github.com/zouzhihao-994/gvm/src/share/classfile"

type ConstantNameAndTypeInfo struct {
	Tag             uint8
	NameIndex       uint16
	DescriptorIndex uint16
	cp              ConstantPool
}

func (constantNameAndTypeInfo *ConstantNameAndTypeInfo) ReadInfo(reader *classfile.ClassReader) {
	constantNameAndTypeInfo.NameIndex = reader.ReadUint16()
	constantNameAndTypeInfo.DescriptorIndex = reader.ReadUint16()
}
