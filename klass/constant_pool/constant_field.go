package constant_pool

import (
	"github.com/zouzhihao-994/gvm/loader"
)

type ConstantFieldInfo struct {
	Tag uint8
	*ConstantPool
	ClassIndex       uint16
	NameAndTypeIndex uint16
	NameAndType      ConstantNameAndTypeInfo
}

func (field *ConstantFieldInfo) ReadInfo(reader *loader.ClassReader) {
	field.ClassIndex = reader.ReadUint16()
	field.NameAndTypeIndex = reader.ReadUint16()
}

func (field *ConstantFieldInfo) ClassName() string {
	return field.GetClassName(field.ClassIndex)
}
func (field *ConstantFieldInfo) NameAndDescriptor() (string, string) {
	return field.GetNameAndType(field.NameAndTypeIndex)
}
