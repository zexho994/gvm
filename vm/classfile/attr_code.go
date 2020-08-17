package classfile

/*
方法属性表
*/
type CodeAttribute struct {
	// 常量池指针
	cp ConstantPool
	// 最大栈
	maxStack uint16
	// 局部变量表大小
	maxLocals uint16
	/*
		方法的内容编译后存放在code表中
		method body after compile
	*/
	code []byte
	// 受检查异常,对应了方法后面throw的部分
	exceptionTable []*ExceptionTableEntry
	// 属性表
	attributes []AttributeInfo
}

/*
受检查异常结构
*/
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
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
