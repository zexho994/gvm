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
// 是一个强类型，可以直接指向的引用
// 可以指向静态方法、实例方法、构造器或者字段
// 指向字段时：实则指向包含字段访问的虚方法，语义上等价于目标自段的getter/setter方法
// 指向方法时：
type ConstantMethodHandleInfo struct {
	Tag            uint8
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func (self ConstantMethodHandleInfo) ReadInfo(reader *classfile.ClassReader) {
	self.ReferenceKind = reader.ReadUint8()
	self.ReferenceIndex = reader.ReadUint16()
}
