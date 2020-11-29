package constant_pool

// tag = CONSTANT_String (8)
//
type ConstantString struct {
	tag    uint8
	strIdx uint16
}
