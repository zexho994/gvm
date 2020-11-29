package constant_pool

type ConstantMethod struct {
	tag            uint8
	classInfoIdx   uint16
	nameAndTypeIdx uint16
}
