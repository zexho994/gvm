package constant_pool

import (
	"github.com/zouzhihao-994/gvm/loader"
)

// 对于非数组类和接口，NameIdx对应的值是类或接口的二进制名称
// 对于数组类，名称会以n个ASCII字符'['开头，随后是数组元素类型的表示
// 如果数组的元素类型是Java原生类型之一，那就以相应的字段描述符来表示
type ConstantClassInfo struct {
	Tag     uint8
	Cp      ConstantPool
	NameIdx uint16
}

func (c *ConstantClassInfo) ReadInfo(reader *loader.ClassReader) {
	c.NameIdx = reader.ReadUint16()
}
func (c *ConstantClassInfo) Name() string {
	return c.Cp.GetUtf8(c.NameIdx)
}
