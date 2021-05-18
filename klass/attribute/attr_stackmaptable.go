package attribute

import "github.com/zouzhihao-994/gvm/loader"

type Attr_StackMapTable struct {
	nameIdx         uint16
	name            string
	attrLen         uint32
	numberOfEntries uint16
	stackMapFrame   []uint16
}

func (attr Attr_StackMapTable) parse(reader *loader.ClassReader) {
	attr.numberOfEntries = reader.ReadUint16()
	// 读取剩余长度的数据，暂时不对stackMapFrame数据进行处理
	reader.ReadBytes(attr.attrLen - 2)
}

func (attr Attr_StackMapTable) Name() string {
	return attr.name
}
