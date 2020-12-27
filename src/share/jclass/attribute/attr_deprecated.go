package attribute

import "github.com/zouzhihao-994/gvm/src/share/classfile"

type Attr_Deprecated struct {
	nameIdx uint16
	name    string
	length  uint32
}

func (attr Attr_Deprecated) parse(reader *classfile.ClassReader) {
	// nothing
}

func (attr Attr_Deprecated) Name() string {
	return ""
}
