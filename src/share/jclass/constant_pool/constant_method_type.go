package constant_pool

import "github.com/zouzhihao-994/gvm/src/share/classfile"

type MethodType struct {
	Tag           uint8
	DescriptorIdx uint16
	Cp            ConstantPool
}

func (method *MethodType) ReadInfo(reader *classfile.ClassReader) {
	method.DescriptorIdx = reader.ReadUint16()
}
