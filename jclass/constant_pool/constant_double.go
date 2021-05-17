package constant_pool

import (
	"github.com/zouzhihao-994/gvm/classfile"
	"math"
)

type ConstantDoubleInfo struct {
	Tag uint8
	val float64
}

func (constantDoubleInfo *ConstantDoubleInfo) ReadInfo(reader *classfile.ClassReader) {
	bytes := reader.ReadUint64()
	constantDoubleInfo.val = math.Float64frombits(bytes)
}

func (constantDoubleInfo *ConstantDoubleInfo) Value() float64 {
	return constantDoubleInfo.val
}