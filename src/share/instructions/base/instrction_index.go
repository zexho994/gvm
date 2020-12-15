package base

type InstructionIndex0 struct {
}

func (i *InstructionIndex0) FetchOperands(reader *MethodCodeReader) {}

type InstructionIndex16 struct {
	Index uint16
}

func (i *InstructionIndex16) FetchOperands(reader *MethodCodeReader) {
	i.Index = reader.ReadUint16()
}
