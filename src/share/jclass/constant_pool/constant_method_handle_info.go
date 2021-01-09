package constant_pool

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
)

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
	Tag            uint8
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func (self ConstantMethodHandleInfo) ReadInfo(reader *classfile.ClassReader) {
	self.ReferenceKind = reader.ReadUint8()
	self.ReferenceIndex = reader.ReadUint16()
}
