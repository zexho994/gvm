package classfile

import "math"

// 常量池中的整数类型
type ConstantIntegerInfo struct{ val int32 }

// 常量池中的浮点数类型
type ConstantFloatInfo struct {
	val float32
}

// 常量池中的长整形类型
type ConstantLongInfo struct {
	val int64
}

type ConstantDoubleInfo struct{ val float64 }

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {

	bytes := reader.readUint64()

	self.val = math.Float64frombits(bytes)
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	// 转化成float32
	self.val = math.Float32frombits(bytes)
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}
