package attribute

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
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
	outerClassIdx        uint16
	innerNameIdx         uint16
	innerClassAccessFlag uint16
}

func (attr *Attr_InnerClasses) parse(reader *classfile.ClassReader) {
	attr.innerClassesNum = reader.ReadUint16()
	attr.innerClasses = make([]innerClass, attr.innerClassesNum)
	for i := 0; i < int(attr.innerClassesNum); i++ {
		attr.innerClasses[i].parse(reader, attr.cp)
	}
}

func (inner innerClass) parse(reader *classfile.ClassReader, pool constant_pool.ConstantPool) {
	inner.innerClassIdx = reader.ReadUint16()
	inner.outerClassIdx = reader.ReadUint16()
	inner.innerNameIdx = reader.ReadUint16()
	inner.innerClassAccessFlag = reader.ReadUint16()
	pool.GetClassName(inner.innerClassIdx)
}

func (attr *Attr_InnerClasses) Name() string {
	return attr.name
}
