package instructions

type CodeReader struct {
	pc   int
	code []byte
}

func (c *CodeReader) Reset(pc int, code []byte) {
	c.pc = pc
	c.code = code
}

func (c *CodeReader) ReadUint8() uint8 {
	i := c.code[c.pc]
	c.pc++
	return i
}

func (c *CodeReader) ReadInt8() int8 {
	return int8(c.ReadUint8())
}

func (c *CodeReader) ReadUint16() uint16 {
	b1 := uint16(c.ReadUint8())
	b2 := uint16(c.ReadUint8())
	return (b1 << 8) | b2
}

func (c *CodeReader) ReadInt16() int16 {
	return int16(c.ReadUint16())
}

func (c *CodeReader) ReadInt32() int32 {
	byte1 := int32(c.ReadUint8())
	byte2 := int32(c.ReadUint8())
	byte3 := int32(c.ReadUint8())
	byte4 := int32(c.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (c *CodeReader) PC() int {
	return c.pc
}
