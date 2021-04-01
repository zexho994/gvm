package constant_pool

import "github.com/zouzhihao-994/gvm/classfile"

// 常量池中的长整形类型
type ConstantLongInfo struct {
	Tag uint8
	val int64
}

func (constantLongInfo *ConstantLongInfo) ReadInfo(reader *classfile.ClassReader) {
	bytes := reader.ReadUint64()
	constantLongInfo.val = int64(bytes)
}

func (constantLongInfo *ConstantLongInfo) Value() int64 {
	return constantLongInfo.val
}
