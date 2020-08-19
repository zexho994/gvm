package base

type BytecodeReader struct {
	// 存放字节码
	code []byte
	// 记录读取到哪一个字节
	pc int
}

func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

/*
uint unsigned 8-bit integers,range:0 through 255
*/
func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

/*
int8 signed 8-bit integers,range:-128 through 127
*/
func (self *BytecodeReader) ReadInt8() int8 { return int8(self.ReadUint8()) }

func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	return (byte1 << 8) | byte2
}

func (self *BytecodeReader) ReadInt16() int16 { return int16(self.ReadUint16()) }

func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}
