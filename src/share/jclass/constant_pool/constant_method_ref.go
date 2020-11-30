package constant_pool

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
)

type ConstantMethod struct {
	Tag            uint8
	Cp             ConstantPool
	ClassIdx       uint16
	NameAndTypeIdx uint16
}

func (ConstantMemberRefInfo *ConstantMethod) ReadInfo(reader *classfile.ClassReader) {
	ConstantMemberRefInfo.ClassIdx = reader.ReadUint16()
	ConstantMemberRefInfo.NameAndTypeIdx = reader.ReadUint16()
}

func (ConstantMemberRefInfo *ConstantMethod) NameAndDescriptor() (string, string) {
	return ConstantMemberRefInfo.Cp.GetNameAndType(ConstantMemberRefInfo.NameAndTypeIdx)
}
