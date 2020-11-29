package classfile

import "github.com/zouzhihao-994/gvm/src/share/jclass"

func ParseToJClass(bytecode []byte) *jclass.JClass {
	reader := &ClassReader{bytecode: bytecode}
	jClass := jclass.JClass{}
	jClass.SetMagic(parseMagic(reader))
	jClass.SetMinorVersion(parseMinorVersion(reader))
	jClass.SetMajorVersion(parseMinorVersion(reader))
	jClass.SetConstantPool(parseConstantPool(reader))
	return &jClass
}

func parseMagic(reader *ClassReader) uint32 {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("[gvm] this file is not support")
	}
	return magic
}

func parseMinorVersion(reader *ClassReader) uint16 {
	return reader.readUint16()
}

func paresMajorVersion(reader *ClassReader) uint16 {
	return reader.readUint16()
}
