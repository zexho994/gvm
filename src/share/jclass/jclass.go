package jclass

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
)

type JClass struct {
	// 魔术
	Magic uint32
	// 次版本
	MinorVersion uint16
	// 主版本
	MajorVersion uint16
	// 常量池
	ConstantPoolCount uint16
	ConstantPool      constant_pool.ConstantPool
	// 类访问标志,表示是类还是接口,public还是private等
	AccessFlags uint16
	// 本类
	ThisClass uint16
	// 父类
	SuperClass uint16
	// 接口
	InterfacesCount uint16
	Interfaces      []uint16
	// 字段表,用于表示接口或者类中声明的变量
	FieldsCount uint16
	Fields      FieldInfo
	// 方法表
	MethodsCount uint16
	Methods      MethodInfo
	// 属性表
	AttributesCount uint16
	Attributes      []AttributeInfo
}

func ParseToJClass(bytecode []byte) *JClass {
	reader := &classfile.ClassReader{Bytecode: bytecode}
	jClass := JClass{}
	jClass.Magic = parseMagic(reader)
	jClass.MinorVersion = parseMinorVersion(reader)
	jClass.MajorVersion = paresMajorVersion(reader)
	jClass.ConstantPoolCount = reader.ReadUint16()
	jClass.ConstantPool = parseConstantPool(jClass.ConstantPoolCount, reader)
	// 类访问符
	jClass.AccessFlags = reader.ReadUint16()
	// 本类
	jClass.ThisClass = reader.ReadUint16()
	// 父类
	jClass.SuperClass = reader.ReadUint16()
	// 接口数量 & 列表
	jClass.InterfacesCount = reader.ReadUint16()
	// 字段数量 & 列表

	// 方法数量 & 列表

	// 属性数量 & 列表

	return &jClass
}

func parseMagic(reader *classfile.ClassReader) uint32 {
	magic := reader.ReadUint32()
	if magic != 0xCAFEBABE {
		panic("[gvm] this file is not support")
	}
	return magic
}

func parseMinorVersion(reader *classfile.ClassReader) uint16 {
	return reader.ReadUint16()
}

func paresMajorVersion(reader *classfile.ClassReader) uint16 {
	return reader.ReadUint16()
}

func parseConstantPool(cpCount uint16, reader *classfile.ClassReader) constant_pool.ConstantPool {
	return constant_pool.ReadConstantPool(cpCount, reader)
}
