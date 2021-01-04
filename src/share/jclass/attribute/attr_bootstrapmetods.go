package attribute

import "github.com/zouzhihao-994/gvm/src/share/classfile"

// 保存invokedynamic指令引用的引导方法限定符号
type BootstrapmethodsAttribute struct {
	// utf8 格式
	nameIdx uint16
	name    string
	// 属性长度，不包括nameIdx和attrLen的6字节
	attrLen    uint32
	methodsNum uint16
	methods    []BootstrapMethod
}

type BootstrapMethod struct {
	// 常量池索引 CONSTANT_MethodHandle_info结构
	methodRef    uint16
	argumentsNum uint16
	arguments    []uint16
}

func (attr *BootstrapmethodsAttribute) parse(reader *classfile.ClassReader) {
	attr.methodsNum = reader.ReadUint16()
	attr.methods = make([]BootstrapMethod, attr.methodsNum)
	for i := 0; i < int(attr.methodsNum); i++ {
		panic("unfinished")
	}
}

func (attr *BootstrapmethodsAttribute) Name() string {
	return attr.name
}
