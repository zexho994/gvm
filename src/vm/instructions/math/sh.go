package math

import (
	"../../rtda"
	"../base"
)

/*
算数左移，无符号左移
*/
type ISHL struct {
	base.NoOperandsInstruction
}

/*
算数右移，无符号右移
*/
type ISHR struct {
	base.NoOperandsInstruction
}

/*
逻辑右移，有符号右移动
*/
type ISHUR struct {
	base.NoOperandsInstruction
}

/*
long整型算数左移
*/
type LSHL struct {
	base.NoOperandsInstruction
}

/*
long整形算数右移
*/
type LSHR struct {
	base.NoOperandsInstruction
}

/*
long整形逻辑右移
*/
type LSHUR struct {
	base.NoOperandsInstruction
}

type DSHL struct {
	base.NoOperandsInstruction
}

type DSHR struct {
	base.NoOperandsInstruction
}

type DSHUR struct {
	base.NoOperandsInstruction
}

type FSHL struct {
	base.NoOperandsInstruction
}

type FSHR struct {
	base.NoOperandsInstruction
}

type FSHUR struct {
	base.NoOperandsInstruction
}

/*
int字符的算数左移,
i << x,i左移x位
*/
func (self ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	s := uint32(v1) & 0x1f
	stack.PushInt(v2 << s)
}

/*
int字符的算数右移
i >> x , i右移动x位
*/
func (self ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	n := uint32(v1) & 0x1f
	stack.PushInt(v2 >> n)
}

/*
int字符的逻辑右移
i >>> x , i右移动x位
go中没有 >>> 运算符，所以先转换成无符号整数，再转换成有符号整数
*/
func (self ISHUR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	n := v1 & 0x1f
	result := int32(uint32(v2) >> n)
	stack.PushInt(result)
}

/*
long 算数左移
*/
func (self LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopLong()

	// long整形64bit = 0x3f
	s := uint32(v1) & 0x3f
	result := v2 << s

	stack.PushLong(result)
}

/*
long 算数右移
算数移动不带符号的，对于整数而言就是补0，对于负数就是补1
*/
func (self LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopLong()

	s := uint32(v1) & 0x3f
	result := v2 >> s

	stack.PushLong(result)
}

/*
long 逻辑右移
*/
func (self LSHUR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopLong()

	n := v1 & 0x3f
	result := int64(uint64(v2) >> n)

	stack.PushLong(result)
}
