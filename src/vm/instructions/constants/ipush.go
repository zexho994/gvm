package constants

import "github.com/zouzhihao-994/gvm/src/vm/instructions/base"
import "github.com/zouzhihao-994/gvm/src/vm/runtime"

type BIPUSH struct{ val int8 } // Push byte

type SIPUSH struct{ val int16 } // Push short

/*
获取一个byte指令
*/
func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

/*
获取一个short指令
*/
func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

/*
将读取到的byte转化成int后推入到栈顶
*/
func (self BIPUSH) Execute(frame *runtime.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

/*
将读取到的short转化成int后推入到栈顶
*/
func (self SIPUSH) Execute(frame *runtime.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
