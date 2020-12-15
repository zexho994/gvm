package constant_pool

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
)

type ConstantMethod struct {
	// contant_method's tag is 10
	Tag uint8
	Cp  ConstantPool
	// represents a class or interface
	classIdx       uint16
	nameAndTypeIdx uint16
}

func (c *ConstantMethod) ReadInfo(reader *classfile.ClassReader) {
	c.classIdx = reader.ReadUint16()
	c.nameAndTypeIdx = reader.ReadUint16()
}

func (c *ConstantMethod) NameAndDescriptor() (string, string) {
	return c.Cp.GetNameAndType(c.nameAndTypeIdx)
}

func (c *ConstantMethod) ClassName() string {
	return c.Cp.GetClassName(c.classIdx)
}
