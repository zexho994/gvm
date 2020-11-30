package constant_pool

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
)

type ConstantClass struct {
	Tag     uint8
	Cp      ConstantPool
	NameIdx uint16
}

func (ConstantClassInfo *ConstantClass) ReadInfo(reader *classfile.ClassReader) {
	ConstantClassInfo.NameIdx = reader.ReadUint16()
}
func (ConstantClassInfo *ConstantClass) Name() string {
	return ConstantClassInfo.Cp.GetUtf8(ConstantClassInfo.NameIdx)
}
