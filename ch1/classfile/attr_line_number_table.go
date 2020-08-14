package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	// 两字节
	lineNumberTableLength := reader.readUint16()
	// 2字节数组
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	// 遍历数组
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc: reader.readUint16(), lineNumber: reader.readUint16(),
		}
	}
}
