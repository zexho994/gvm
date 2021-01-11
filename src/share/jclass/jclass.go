package jclass

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	"github.com/zouzhihao-994/gvm/src/share/jclass/attribute"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
)

// JClass 是将字节码进行翻译后的结构，也就是仅执行加载步骤后的形态
// 这时候还没有进行链接、初始化等步骤，所以 JClass 中存储的还是符号引用，非直接引用
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
	ThisClassIdx uint16
	// 父类
	SuperClassIdx uint16
	// 接口
	InterfacesCount uint16
	Interfaces      []uint16
	// 字段表,用于表示接口或者类中声明的变量
	FieldsCount uint16
	Fields      Fields
	// 方法表
	MethodsCount uint16
	Methods      Methods
	// 属性表
	AttributesCount uint16
	Attributes      attribute.AttributesInfo
}

// 可以理解为类加载阶段中的<加载>步骤
func ParseToJClass(bytecode []byte) *JClass {
	reader := &classfile.ClassReader{Bytecode: bytecode}
	jClass := JClass{}
	// CAFEBABY
	jClass.Magic = parseMagic(reader)
	// jdk version
	jClass.MinorVersion = parseMinorVersion(reader)
	jClass.MajorVersion = paresMajorVersion(reader)
	// 常量池
	jClass.ConstantPoolCount = reader.ReadUint16()
	jClass.ConstantPool = parseConstantPool(jClass.ConstantPoolCount, reader)
	// 类访问符
	jClass.AccessFlags = reader.ReadUint16()
	// 本类
	jClass.ThisClassIdx = reader.ReadUint16()
	// 父类
	jClass.SuperClassIdx = reader.ReadUint16()
	// 接口数量 & 列表
	jClass.InterfacesCount = reader.ReadUint16()
	jClass.Interfaces = reader.ReadUint16Array(jClass.InterfacesCount)
	// 字段数量 & 列表
	jClass.FieldsCount = reader.ReadUint16()
	jClass.Fields = parseFields(jClass.FieldsCount, reader, jClass.ConstantPool)
	// 方法数量 & 列表
	jClass.MethodsCount = reader.ReadUint16()
	jClass.Methods = parseMethod(jClass.MethodsCount, reader, jClass.ConstantPool)
	// 属性数量 & 列表
	jClass.AttributesCount = reader.ReadUint16()
	jClass.Attributes = attribute.ParseAttributes(jClass.AttributesCount, reader, jClass.ConstantPool)

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
