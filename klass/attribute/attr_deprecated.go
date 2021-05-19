package attribute

import "github.com/zouzhihao-994/gvm/loader"

type AttrDeprecated struct {
	nameIdx uint16
	name    string
	attrLen uint32
}

func (attr AttrDeprecated) parse(reader *loader.ClassReader) {
	// nothing
}

func (attr AttrDeprecated) Name() string {
	return ""
}
