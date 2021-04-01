package constant_pool

import (
	"github.com/zouzhihao-994/gvm/classfile"
)

type ConstantInterfaceMethodInfo struct {
	Tag            uint8
	Cp             ConstantPool
	ClassIdx       uint16
	NameAndTypeIdx uint16
}

/*
读取数据
*/
func (im *ConstantInterfaceMethodInfo) ReadInfo(reader *classfile.ClassReader) {
	im.ClassIdx = reader.ReadUint16()
	im.NameAndTypeIdx = reader.ReadUint16()
}

/*
获取类名
*/
func (im *ConstantInterfaceMethodInfo) ClassName() string {
	return im.Cp.GetClassName(im.ClassIdx)
}

/*
获取描述符
*/
func (im *ConstantInterfaceMethodInfo) NameAndDescriptor() (string, string) {
	return im.Cp.GetNameAndType(im.NameAndTypeIdx)
}
