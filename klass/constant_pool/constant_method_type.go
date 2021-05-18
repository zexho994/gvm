package constant_pool

import "github.com/zouzhihao-994/gvm/loader"

type ConstantMethodTypeInfo struct {
	Tag           uint8
	DescriptorIdx uint16
	Cp            ConstantPool
}

func (method *ConstantMethodTypeInfo) ReadInfo(reader *loader.ClassReader) {
	method.DescriptorIdx = reader.ReadUint16()
}