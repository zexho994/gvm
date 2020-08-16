package classfile

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

func readConstantInvokeDynamicInfo(reader *ClassReader) ConstantInvokeDynamicInfo {
	return ConstantInvokeDynamicInfo{
		BootstrapMethodAttrIndex: reader.readUint16(),
		NameAndTypeIndex:         reader.readUint16(),
	}
}

func (self ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.NameAndTypeIndex = reader.readUint16()
	self.BootstrapMethodAttrIndex = reader.readUint16()
}

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func readConstantMethodHandleInfo(reader *ClassReader) ConstantMethodHandleInfo {
	return ConstantMethodHandleInfo{
		ReferenceKind:  reader.readUint8(),
		ReferenceIndex: reader.readUint16(),
	}
}

func (self ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	self.ReferenceKind = reader.readUint8()
	self.ReferenceIndex = reader.readUint16()
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodTypeInfo struct {
	DescriptorIndex uint16
}

func readConstantMethodTypeInfo(reader *ClassReader) ConstantMethodTypeInfo {
	return ConstantMethodTypeInfo{
		DescriptorIndex: reader.readUint16(),
	}
}

func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.DescriptorIndex = reader.readUint16()
}
