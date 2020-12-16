package attribute

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
)

type Attr_Code struct {
	NameIdx uint16
	name    string
	AttrLen uint32
	cp      constant_pool.ConstantPool
	// 方法的操作数栈在任何时间点的最大深度
	// 最大深度值在编译期就可以确定
	MaxStack uint16
	// 局部变量表大小，包括方法的参数
	MaxLocals uint16
	codeLen   uint32
	code      []byte
	// 异常表
	ExceptionTable []*ExceptionTable
	// 属性表
	attrCount uint16
	attrInfo  AttributeInfos
}

// 异常表
type ExceptionTable struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (c *Attr_Code) parse(reader *classfile.ClassReader) {
	c.MaxStack = reader.ReadUint16()
	c.MaxLocals = reader.ReadUint16()
	c.codeLen = reader.ReadUint32()
	c.code = reader.ReadBytes(c.codeLen)
	c.ExceptionTable = parseExceptionTable(reader)
	c.attrCount = reader.ReadUint16()
	c.attrInfo = ParseAttributes(c.attrCount, reader, c.cp)
}

func parseExceptionTable(reader *classfile.ClassReader) []*ExceptionTable {
	tableLen := reader.ReadUint16()
	table := make([]*ExceptionTable, tableLen)
	for i := range table {
		table[i] = &ExceptionTable{
			startPc:   reader.ReadUint16(),
			endPc:     reader.ReadUint16(),
			handlerPc: reader.ReadUint16(),
			catchType: reader.ReadUint16(),
		}
	}
	return table
}

func CreateCodeAttr(maxStack, maxLocal uint16, code []byte) *Attr_Code {
	return &Attr_Code{
		MaxStack:  maxStack,
		MaxLocals: maxLocal,
		code:      code,
	}
}

func (c Attr_Code) Name() string {
	return c.name
}

func (c Attr_Code) Code() []byte {
	return c.code
}
