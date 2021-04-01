package constant_pool

import (
	"github.com/zouzhihao-994/gvm/classfile"
)

// Tag = CONSTANT_String (8)
//
type ConstantStringInfo struct {
	Tag    uint8
	Cp     ConstantPool
	strIdx uint16
}

/*
读取字符串的常量池索引
*/
func (constantStringInfo *ConstantStringInfo) ReadInfo(reader *classfile.ClassReader) {
	constantStringInfo.strIdx = reader.ReadUint16()
}

/*
输出常量池中字符串的值
*/
func (constantStringInfo *ConstantStringInfo) String() string {
	return constantStringInfo.Cp.GetUtf8(constantStringInfo.strIdx)
}
