package classfile

import "math"

// 常量池中的整数类型
type ConstantIntegerInfo struct {
	tag uint8
	val int32
}

// 常量池中的浮点数类型
type ConstantFloatInfo struct {
	tag uint8
	val float32
}

// 常量池中的长整形类型
type ConstantLongInfo struct {
	tag uint8
	val int64
}

type ConstantDoubleInfo struct {
	tag uint8
	val float64
}

func (constantDoubleInfo *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	constantDoubleInfo.val = math.Float64frombits(bytes)
}

func (constantIntegerInfo *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	constantIntegerInfo.val = int32(bytes)
}

func (constantFloatInfo *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	// 转化成float32
	constantFloatInfo.val = math.Float32frombits(bytes)
}

func (constantLongInfo *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	constantLongInfo.val = int64(bytes)
}

func (constantDoubleInfo *ConstantDoubleInfo) Value() float64 {
	return constantDoubleInfo.val
}

func (constantLongInfo *ConstantLongInfo) Value() int64 {
	return constantLongInfo.val
}

func (constantFloatInfo *ConstantFloatInfo) Value() float32 {
	return constantFloatInfo.val
}

func (constantIntegerInfo *ConstantIntegerInfo) Value() int32 {
	return constantIntegerInfo.val
}
