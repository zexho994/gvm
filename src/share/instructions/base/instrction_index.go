package base

type InstructionIndex0 struct {
	Index uint16
}

func (i *InstructionIndex0) FetchOperands(reader *MethodCodeReader) {}

type InstructionIndex8 struct {
	Index uint8
}

func (i *InstructionIndex8) FetchOperands(reader *MethodCodeReader) {
	i.Index = reader.ReadUint8()
}

type InstructionIndex16 struct {
	Index uint16
}

func (i *InstructionIndex16) FetchOperands(reader *MethodCodeReader) {
	i.Index = reader.ReadUint16()
}
