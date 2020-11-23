package classfile

/*
tag
常量一般都由两个部分组成,tag和信息
tag用来区分常量类型
*/
const (
	// UTF-8编码的字符串
	ConstantUtf8 = 1
	// 整形字面量
	ConstantInteger = 3
	// 浮点数字面量
	ConstantFloat = 4
	// 长整形字面量
	ConstantLong = 5
	// 双精度浮点型字面量
	ConstantDouble = 6
	// 类或者接口的符号引用
	ConstantClass = 7
	// 字符串类型字面量
	ConstantString = 8
	// 字段的符号引用
	ConstantFieldRef = 9
	// 类中方法的符号引用
	ConstantMethodRef = 10
	// 接口中方法的符号引用
	ConstantInterfaceMethodRef = 11
	// 字段或者方法的部分符号引用
	ConstantNameAndType = 12
	// 方法句柄
	ConstantMethodHandle = 15
	// 方法类型
	ConstantMethodType = 16
	// 动态方法调用点
	ConstantInvokeDynamic = 18
)
