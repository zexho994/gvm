package attribute

import "github.com/zouzhihao-994/gvm/loader"

type AttrSynthetic struct {
	nameIdx uint16
	name    string
	attrLen uint32
}

func (attr AttrSynthetic) parse(reader *loader.ClassReader) {
	// nothing
}

func (attr AttrSynthetic) Name() string {
	return ""
}
