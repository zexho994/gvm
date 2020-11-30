package constant_pool

import "github.com/zouzhihao-994/gvm/src/share/classfile"

// 常量池中的长整形类型
type ConstantLong struct {
	Tag uint8
	val int64
}

func (constantLongInfo *ConstantLong) ReadInfo(reader *classfile.ClassReader) {
	bytes := reader.ReadUint64()
	constantLongInfo.val = int64(bytes)
}

func (constantLongInfo *ConstantLong) Value() int64 {
	return constantLongInfo.val
}
