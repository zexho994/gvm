package classfile

type ConstantIntegerInfo struct{ val int32 }

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}
