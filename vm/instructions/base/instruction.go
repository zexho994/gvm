package base

import "../../rtda"

/*
指令接口，指令主要两个功能：
1 从字节码中获取操作数
2 执行操作指令
*/
type Instruction interface {
	// 从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	// 操作指令
	Execute(frame *rtda.Frame)
}

/*
无操作指令 结构体
提供给其他相同没有操作的指令结构体调用
*/
type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

/*
分支指令结构体
一般用于if else，for等条件语句中
*/
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

/*
根据索引获取局部变量表
*/
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

/*
16位长度的指令
*/
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
