package constant_pool

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
)

type ConstantInterfaceMethod struct {
	Tag            uint8
	Cp             ConstantPool
	ClassIdx       uint16
	NameAndTypeIdx uint16
}

/*
读取数据
*/
func (im *ConstantInterfaceMethod) ReadInfo(reader *classfile.ClassReader) {
	im.ClassIdx = reader.ReadUint16()
	im.NameAndTypeIdx = reader.ReadUint16()
}

/*
获取类名
*/
func (im *ConstantInterfaceMethod) ClassName() string {
	return im.Cp.GetClassName(im.ClassIdx)
}

/*
获取描述符
*/
func (im *ConstantInterfaceMethod) NameAndDescriptor() (string, string) {
	return im.Cp.GetNameAndType(im.NameAndTypeIdx)
}
