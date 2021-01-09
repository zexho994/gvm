package constant_pool

import "github.com/zouzhihao-994/gvm/src/share/classfile"

/*
符号引用的指针都会指向常量池中的索引
索引对应的数据用本utf8结构显示
*/
type ConstantUtf8 struct {
	Tag uint8
	Str string
}

func (utf8 *ConstantUtf8) ReadInfo(reader *classfile.ClassReader) {
	length := uint32(reader.ReadUint16())
	bytes := reader.ReadBytes(length)
	utf8.Str = decodeMUTF8(bytes)
}

/*
完整的在源码中补充
*/
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
