package extended

import (
	"../../instructions/base"
	"../../instructions/loads"
	"../../rtda"
)

/*
利用WIDE指令可以根据对应的指令进行扩展
<p>
扩展指令用于一般的指令无法完成的情况下
例如局部变量表默认为256bit，所以一个字节的索引就可以搜索到所有的数据
如果局部变量表的大小大于256bit的情况下，就需要扩展索引的大小
<p>
*/
type WIDE struct {
	modifiedInstruction base.Instruction
}

/*

 */
func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst // iload
	case 0x16:
		inst := &loads.LLOAD{} // lload
		inst.Index = uint(reader.ReadInt32())
		self.modifiedInstruction = inst
	case 0x17:

	case 0x18:
		// dload
	case 0x19:
		// aload
	case 0x36:
		// istore
	case 0x37:
		// lstore
	case 0x38:
		// fstore
	case 0x39:
		// dstore
	case 0x3a:
		// astore
	case 0x84:
		// iinc
	case 0xa9: // ret
		panic("Unsupported opcode: 0xa9!")
	}
}

func (self *WIDE) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}
