package attribute

// 保存invokedynamic指令引用的引导方法限定符号
type BootstrapmethodsAttribute struct {
	// utf8 格式
	nameIdx uint16
	// 属性长度，不包括nameIdx和attrLen的6字节
	attrLen    uint32
	methodsNum uint16
}

type BootstrapMethod struct {
	// 常量池索引 CONSTANT_MethodHandle_info结构
	methodRef    uint16
	argumentsNum uint16
	arguments    []uint16
}
