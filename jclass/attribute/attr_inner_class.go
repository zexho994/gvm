package attribute

import (
	"github.com/zouzhihao-994/gvm/classfile"
	"github.com/zouzhihao-994/gvm/jclass/constant_pool"
)

type Attr_InnerClasses struct {
	nameIdx         uint16
	name            string
	attrLen         uint32
	cp              constant_pool.ConstantPool
	innerClassesNum uint16
	innerClasses    []innerClass
}

type innerClass struct {
	// Constant_Class_Info type
	innerClassIdx        uint16
	innerClassStr        string
	outerClassIdx        uint16
	outerClassStr        string
	innerNameIdx         uint16
	innerNameStr         string
	innerClassAccessFlag uint16
}

func (attr *Attr_InnerClasses) parse(reader *classfile.ClassReader) {
	attr.innerClassesNum = reader.ReadUint16()
	attr.innerClasses = make([]innerClass, attr.innerClassesNum)
	for i := 0; i < int(attr.innerClassesNum); i++ {
		attr.innerClasses[i].parse(reader, attr.cp)
	}
}

func (inner *innerClass) parse(reader *classfile.ClassReader, pool constant_pool.ConstantPool) {
	inner.innerClassIdx = reader.ReadUint16()
	inner.outerClassIdx = reader.ReadUint16()
	inner.innerNameIdx = reader.ReadUint16()
	inner.innerClassAccessFlag = reader.ReadUint16()
	inner.innerClassStr = pool.GetClassName(inner.innerClassIdx)
	if inner.innerNameIdx != 0 {
		inner.innerNameStr = pool.GetUtf8(inner.innerNameIdx)
	}
	if inner.outerClassIdx != 0 {
		inner.outerClassStr = pool.GetClassName(inner.outerClassIdx)
	}
}

func (inner *innerClass) innerClasses() *innerClass {
	return inner
}

func (attr *Attr_InnerClasses) Name() string {
	return attr.name
}
