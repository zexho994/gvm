package attribute

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
)

// LineNumberTable属性是可选的变长属性，位于Code结构的属性表中
// 用于调试器确定源文件中由给定行所表示的内容，对应code[]数组中的哪一个部分
// LineNumberTable属性可以按照任意顺序出现
// 可以有不止一个LineNumberTable属性对应于源文件中的同一行.也就是说多个LineNumberTable属性可以合起来表示源文件中的某行代码
// 属性于源文件的代码之间不必有一一对应关系
type Attr_LineNumberTable struct {
	nameIdx uint16
	// 常量池中 nameIdx 对应的UTF8字符
	name            string
	attrLen         uint32
	tableCount      uint16
	lineNumberTable []lineNumberTable
	cp              constant_pool.ConstantPool
}

type lineNumberTable struct {
	// 必须是code[]数组的一个索引，code[] 在该索引出的指令码
	startPC uint16
	// 与源文件中对应
	lineNumber uint16
}

func (attr *Attr_LineNumberTable) parse(reader *classfile.ClassReader) {
	attr.tableCount = reader.ReadUint16()
	table := make([]lineNumberTable, attr.tableCount)
	for i := range table {
		lnt := lineNumberTable{}
		lnt.startPC = reader.ReadUint16()
		lnt.lineNumber = reader.ReadUint16()
		table[i] = lnt
	}
	attr.lineNumberTable = table
}

func (attr Attr_LineNumberTable) Name() string {
	return attr.name
}
