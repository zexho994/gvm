package attribute

import "github.com/zouzhihao-994/gvm/loader"

type Attr_SourceFile struct {
	nameIdx uint16
	name    string
	attrLen uint32
	// 对应常量池中的 UTF8 类型索引，表示class文件源文件的名字
	sourceIdx uint16
}

func (attr *Attr_SourceFile) parse(reader *loader.ClassReader) {
	attr.sourceIdx = reader.ReadUint16()
}

func (attr Attr_SourceFile) Name() string {
	return attr.name
}
