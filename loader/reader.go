package loader

type ClassReader struct {
	Bytecode []byte
}

func NewClassReader(b []byte) *ClassReader {
	return &ClassReader{Bytecode: b}
}

func (reader *ClassReader) ReadUint8() uint8 {
	data := reader.Bytecode[0]
	reader.Bytecode = reader.Bytecode[1:]
	return data
}

func (reader *ClassReader) ReadInt8() int8 {
	return int8(reader.ReadUint8())
}

func (reader *ClassReader) ReadUint16() uint16 {
	_ = reader.Bytecode[1]
	val := uint16(reader.Bytecode[1]) | uint16(reader.Bytecode[0])<<8
	reader.Bytecode = reader.Bytecode[2:]
	return val
}

func (reader *ClassReader) ReadInt16() int16 {
	return int16(reader.ReadUint16())
}

func (reader *ClassReader) ReadUint16Array(count uint16) []uint16 {
	data := make([]uint16, count)
	for i := range data {
		data[i] = reader.ReadUint16()
	}
	return data
}

func (reader *ClassReader) ReadUint32() uint32 {
	_ = reader.Bytecode[3]
	val := uint32(reader.Bytecode[3]) |
		uint32(reader.Bytecode[2])<<8 |
		uint32(reader.Bytecode[1])<<16 |
		uint32(reader.Bytecode[0])<<24
	reader.Bytecode = reader.Bytecode[4:]
	return val
}

func (reader *ClassReader) ReadUint64() uint64 {
	_ = reader.Bytecode[7]
	val := uint64(reader.Bytecode[7]) |
		uint64(reader.Bytecode[6])<<8 |
		uint64(reader.Bytecode[5])<<16 |
		uint64(reader.Bytecode[4])<<24 |
		uint64(reader.Bytecode[3])<<32 |
		uint64(reader.Bytecode[2])<<40 |
		uint64(reader.Bytecode[1])<<48 |
		uint64(reader.Bytecode[0])<<56
	reader.Bytecode = reader.Bytecode[8:]
	return val
}

// ReadBytes 读取length长度的数据
func (reader *ClassReader) ReadBytes(length uint32) []byte {
	val := reader.Bytecode[:length]
	reader.Bytecode = reader.Bytecode[length:]
	return val
}
