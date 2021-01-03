package constant_pool

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
)

type ConstantFieldRefInfo struct {
	Tag              uint8
	Cp               ConstantPool
	ClassIndex       uint16
	NameAndTypeIndex uint16
	NameAndType      ConstantNameAndType
}

func (field *ConstantFieldRefInfo) ReadInfo(reader *classfile.ClassReader) {
	field.ClassIndex = reader.ReadUint16()
	field.NameAndTypeIndex = reader.ReadUint16()
}

func (field *ConstantFieldRefInfo) ClassName() string {
	return field.Cp.GetClassName(field.ClassIndex)
}
func (field *ConstantFieldRefInfo) NameAndDescriptor() (string, string) {
	return field.Cp.GetNameAndType(field.NameAndTypeIndex)
}
