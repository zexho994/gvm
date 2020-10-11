package classfile

/*
tag
常量一般都由两个部分组成,tag和信息
tag用来区分常量类型
*/
const (
	// 类或者接口的符号引用
	CONSTANT_Class = 7
	// 字段的符号引用
	CONSTANT_Fieldref = 9
	// 类中方法的符号引用
	CONSTANT_Methodref = 10
	// 接口中方法的符号引用
	CONSTANT_InterfaceMethodref = 11
	// 字符串类型字面量
	CONSTANT_String = 8
	// 整形字面量
	CONSTANT_Integer = 3
	// 浮点数字面量
	CONSTANT_Float = 4
	// 长整形字面量
	CONSTANT_Long = 5
	// 双精度浮点型字面量
	CONSTANT_Double = 6
	// 字段或者方法的部分符号引用
	CONSTANT_NameAndType = 12
	// UTF-8编码的字符串
	CONSTANT_Utf8 = 1
	// 方法句柄
	CONSTANT_MethodHandle = 15
	// 方法类型
	CONSTANT_MethodType = 16
	// 动态方法调用点
	CONSTANT_InvokeDynamic = 18
)
