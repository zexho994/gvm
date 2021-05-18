package attribute

import "github.com/zouzhihao-994/gvm/classloader"

type Attr_Deprecated struct {
	nameIdx uint16
	name    string
	attrLen uint32
}

func (attr Attr_Deprecated) parse(reader *classloader.ClassReader) {
	// nothing
}

func (attr Attr_Deprecated) Name() string {
	return ""
}
