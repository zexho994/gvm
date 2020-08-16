package classfile

type LocalVariableTableAttribute struct{ localvariabletable []*LineNumberTableEntry }

type LocalVariableTableAttributeEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	local := reader.readUint16()
	self.localvariabletable = make([]*LineNumberTableEntry, local)
	for i := range self.localvariabletable {
		self.localvariabletable[i] = &LineNumberTableEntry{
			startPc: reader.readUint16(), lineNumber: reader.readUint16(),
		}
	}
}
