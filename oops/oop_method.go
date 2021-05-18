package oops

type OopMethod struct {
	access       uint16
	name         string
	descriptor   string
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func (m OopMethod) Access() uint16 {
	return m.access
}

func (m OopMethod) Name() string {
	return m.name
}

func (m OopMethod) Descriptor() string {
	return m.descriptor
}

func (m OopMethod) MaxStack() uint {
	return m.maxStack
}

func (m OopMethod) MaxLocals() uint {
	return m.maxLocals
}

func (m OopMethod) Code() []byte {
	return m.code
}

func (m OopMethod) ArgSlotCount() uint {
	return m.argSlotCount
}
