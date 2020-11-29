package classfile

import (
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
)

func parseConstantPool(reader *ClassReader) jclass.Constant {
	cpCount := reader.ReadUint16()
	cp := make(jclass.ConstantPool, cpCount)
	for i := uint16(1); i < cpCount; i++ {
		tag := reader.ReadUint8()
		c := newConstantInfo(tag,cp)
		cp[i] = c
		switch cp[i].Tag(type) {
		// long or double need two bytes
			case *ConstantLongInfo, *ConstantDoubleInfo:
				i++
			}
		}
	return cp
}

