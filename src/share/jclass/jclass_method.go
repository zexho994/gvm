package jclass

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	"github.com/zouzhihao-994/gvm/src/share/jclass/attribute"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
)

type Methods []MethodInfo

type MethodInfo struct {
	accessFlag    uint16
	nameIdx       uint16
	descriptorIdx uint16
	attrCount     uint16
	cp            constant_pool.ConstantPool
	attribute     attribute.AttributeInfo
}

func parseMethod(count uint16, reader *classfile.ClassReader, pool constant_pool.ConstantPool) Methods {
	methods := make([]MethodInfo, count)
	return methods
}
