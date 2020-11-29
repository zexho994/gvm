package constant_pool

type ConstantInterface struct {
	tag            uint8
	classIdx       uint16
	nameAndTypeIdx uint16
}
