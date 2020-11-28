package classfile

type ClassReader struct {
	bytecode []byte
}

func (reader *ClassReader) readUint8() uint8 {
	data := reader.bytecode[0]
	reader.bytecode = reader.bytecode[1:]
	return data
}

func (reader *ClassReader) readUint16() uint16 {
	_ = reader.bytecode[1]
	val := uint16(reader.bytecode[1]) | uint16(reader.bytecode[0])<<8
	reader.bytecode = reader.bytecode[2:]
	return val
}

func (reader *ClassReader) readUint32() uint32 {
	_ = reader.bytecode[3]
	val := uint32(reader.bytecode[3]) |
		uint32(reader.bytecode[2])<<8 |
		uint32(reader.bytecode[1])<<16 |
		uint32(reader.bytecode[0])<<24
	reader.bytecode = reader.bytecode[4:]
	return val
}

func (reader *ClassReader) readUint64() uint64 {
	_ = reader.bytecode[7]
	val := uint64(reader.bytecode[7]) |
		uint64(reader.bytecode[6])<<8 |
		uint64(reader.bytecode[5])<<16 |
		uint64(reader.bytecode[4])<<24 |
		uint64(reader.bytecode[3])<<32 |
		uint64(reader.bytecode[2])<<40 |
		uint64(reader.bytecode[1])<<48 |
		uint64(reader.bytecode[0])<<56
	reader.bytecode = reader.bytecode[8:]
	return val
}

// 读取length长度的数据
func (reader *ClassReader) readBytes(length uint32) []byte {
	val := reader.bytecode[:length]
	reader.bytecode = reader.bytecode[length:]
	return val
}
