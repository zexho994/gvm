package constant_pool

import "github.com/zouzhihao-994/gvm/classloader"

type ConstantMethodTypeInfo struct {
	Tag           uint8
	DescriptorIdx uint16
	Cp            ConstantPool
}

func (method *ConstantMethodTypeInfo) ReadInfo(reader *classloader.ClassReader) {
	method.DescriptorIdx = reader.ReadUint16()
}
