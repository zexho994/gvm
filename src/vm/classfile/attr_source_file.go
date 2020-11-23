package classfile

/*
source_file属性只会出现在classfile目录中
*/
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (sourceFileAttribute *SourceFileAttribute) readInfo(reader *ClassReader) {
	sourceFileAttribute.sourceFileIndex = reader.readUint16()
}

func (sourceFileAttribute *SourceFileAttribute) FileName() string {
	return sourceFileAttribute.cp.getUtf8(sourceFileAttribute.sourceFileIndex)
}
