package base

type MethodCodeReader struct {
	code []byte
	pc   int
}

func (r *MethodCodeReader) PC() int {
	return r.pc
}

/*
重新赋值
*/
func (r *MethodCodeReader) Reset(code []byte, pc int) {
	r.code = code
	r.pc = pc
}

func (r *MethodCodeReader) ReadOpenCdoe() uint8 {
	i := r.code[r.pc]
	r.pc++
	return i
}

func (r *MethodCodeReader) ReadUint8() uint8 {
	i := r.code[r.pc]
	r.pc++
	return i
}

func (r *MethodCodeReader) ReadUint16() uint16 {
	byte1 := uint16(r.ReadUint8())
	byte2 := uint16(r.ReadUint8())
	return (byte1 << 8) | byte2
}
