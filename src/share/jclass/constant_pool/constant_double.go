package constant_pool

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	"math"
)

type ConstantDouble struct {
	Tag uint8
	val float64
}

func (constantDoubleInfo *ConstantDouble) ReadInfo(reader *classfile.ClassReader) {
	bytes := reader.ReadUint64()
	constantDoubleInfo.val = math.Float64frombits(bytes)
}

func (constantDoubleInfo *ConstantDouble) Value() float64 {
	return constantDoubleInfo.val
}
