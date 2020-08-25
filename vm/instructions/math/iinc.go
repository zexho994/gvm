package math

import (
	"../../instructions/base"
	"../../rtda"
)

type IINC struct {
	// 局部变量表下标
	Index uint

	// 操作数
	Const int32
}

func (self IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadInt8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}
