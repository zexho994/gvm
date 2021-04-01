package oops

type Oop_Method struct {
	access       uint16
	name         string
	descriptor   string
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func (m Oop_Method) Access() uint16 {
	return m.access
}

func (m Oop_Method) Name() string {
	return m.name
}

func (m Oop_Method) Descriptor() string {
	return m.descriptor
}

func (m Oop_Method) MaxStack() uint {
	return m.maxStack
}

func (m Oop_Method) MaxLocals() uint {
	return m.maxLocals
}

func (m Oop_Method) Code() []byte {
	return m.code
}

func (m Oop_Method) ArgSlotCount() uint {
	return m.argSlotCount
}
