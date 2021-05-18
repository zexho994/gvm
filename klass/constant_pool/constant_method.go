package constant_pool

import (
	"github.com/zouzhihao-994/gvm/loader"
)

// ConstantMethodInfo 对应常量池中的 Methodref
type ConstantMethodInfo struct {
	// contant_method's tag is 10
	Tag uint8
	Cp  ConstantPool
	// represents a class or interface
	classIdx       uint16
	nameAndTypeIdx uint16
}

func (c *ConstantMethodInfo) ReadInfo(reader *loader.ClassReader) {
	c.classIdx = reader.ReadUint16()
	c.nameAndTypeIdx = reader.ReadUint16()
}

func (c *ConstantMethodInfo) NameAndDescriptor() (string, string) {
	return c.Cp.GetNameAndType(c.nameAndTypeIdx)
}

func (c *ConstantMethodInfo) ClassName() string {
	return c.Cp.GetClassName(c.classIdx)
}