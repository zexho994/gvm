package classfile

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (unparsedAttribute *UnparsedAttribute) readInfo(reader *ClassReader) {
	unparsedAttribute.info = reader.readBytes(unparsedAttribute.length)
}
