package constant_pool

import (
	"github.com/zouzhihao-994/gvm/classloader"
	"math"
)

// 常量池中的浮点数类型
type ConstantFloatInfo struct {
	Tag uint8
	val float32
}

func (constantFloatInfo *ConstantFloatInfo) ReadInfo(reader *classloader.ClassReader) {
	bytes := reader.ReadUint32()
	// 转化成float32
	constantFloatInfo.val = math.Float32frombits(bytes)
}

func (constantFloatInfo *ConstantFloatInfo) Value() float32 {
	return constantFloatInfo.val
}
