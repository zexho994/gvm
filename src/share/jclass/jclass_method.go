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
	attribute     []attribute.AttributeInfo
	cp            constant_pool.ConstantPool
}

// 解析方法表
func parseMethod(count uint16, reader *classfile.ClassReader, pool constant_pool.ConstantPool) Methods {
	methods := make([]MethodInfo, count)
	for i := range methods {
		method := MethodInfo{}
		method.cp = pool
		method.accessFlag = reader.ReadUint16()
		method.nameIdx = reader.ReadUint16()
		method.descriptorIdx = reader.ReadUint16()
		method.attrCount = reader.ReadUint16()
		// 解析方法表中的属性表字段
		method.attribute = attribute.ParseAttributes(method.attrCount, reader, pool)
		methods[i] = method
	}
	return methods
}
