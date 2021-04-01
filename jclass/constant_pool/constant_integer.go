package constant_pool

import "github.com/zouzhihao-994/gvm/classfile"

// 常量池中的整数类型
type ConstantIntegerInfo struct {
	Tag uint8
	val int32
}

func (constantIntegerInfo *ConstantIntegerInfo) ReadInfo(reader *classfile.ClassReader) {
	bytes := reader.ReadUint32()
	constantIntegerInfo.val = int32(bytes)
}

func (constantIntegerInfo *ConstantIntegerInfo) Value() int32 {
	return constantIntegerInfo.val
}
