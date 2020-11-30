package constant_pool

import "github.com/zouzhihao-994/gvm/src/share/classfile"

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamic struct {
	Tag                      uint8
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

func readConstantInvokeDynamicInfo(reader *classfile.ClassReader) ConstantInvokeDynamic {
	return ConstantInvokeDynamic{
		BootstrapMethodAttrIndex: reader.ReadUint16(),
		NameAndTypeIndex:         reader.ReadUint16(),
	}
}

func (invoke ConstantInvokeDynamic) ReadInfo(reader *classfile.ClassReader) {
	invoke.NameAndTypeIndex = reader.ReadUint16()
	invoke.BootstrapMethodAttrIndex = reader.ReadUint16()
}
