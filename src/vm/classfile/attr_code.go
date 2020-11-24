package classfile

// 方法属性表
type CodeAttribute struct {
	// 常量池指针
	cp ConstantPool

	// 最大栈
	maxStack uint16

	// 局部变量表大小
	maxLocals uint16

	//	方法的内容编译后存放在code表中
	//	method body after compile
	//	内容就是iconst_0,istore_1等
	code []byte

	// 受检查异常,对应了方法后面throw的部分
	exceptionTable []*ExceptionTableEntry

	// 属性表
	attributesCount uint16
	attributes      []AttributeInfo
}

// 受检查异常结构
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (codeAttribute *CodeAttribute) readInfo(reader *ClassReader) {
	codeAttribute.maxStack = reader.readUint16()
	codeAttribute.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	codeAttribute.code = reader.readBytes(codeLength)
	codeAttribute.exceptionTable = readExceptionTable(reader)
	codeAttribute.attributesCount = reader.readUint16()
	codeAttribute.attributes = readAttributes(codeAttribute.attributesCount, reader, codeAttribute.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc: reader.readUint16(), endPc: reader.readUint16(), handlerPc: reader.readUint16(), catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

/*
操作数栈最大大小
*/
func (codeAttribute CodeAttribute) MaxStack() uint {
	return uint(codeAttribute.maxStack)
}

/*
局部变量表最大值
*/
func (codeAttribute CodeAttribute) MaxLocals() uint {
	return uint(codeAttribute.maxLocals)
}

func (codeAttribute CodeAttribute) Code() []byte {
	return codeAttribute.code
}
