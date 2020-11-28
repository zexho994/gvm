package jclass

type JClass struct {
	// 魔术
	magic uint32
	// 次版本
	minorVersion uint16
	// 主版本
	majorVersion uint16
	// 常量池
	constantPoolCount uint16
	constantPool      ConstantPool
	// 类访问标志,表示是类还是接口,public还是private等
	accessFlags uint16
	// 本类
	thisClass uint16
	// 父类
	superClass uint16
	// 接口
	interfacesCount uint16
	interfaces      []uint16
	// 字段表,用于表示接口或者类中声明的变量
	fieldsCount uint16
	fields      FieldInfo
	// 方法表
	methodsCount uint16
	methods      MethodInfo
	// 属性表
	attributesCount uint16
	attributes      []AttributeInfo
}

func (c *JClass) SetMagic(magic uint32) {
	c.magic = magic
}

func (c *JClass) SetMinorVersion(v uint16) {
	c.minorVersion = v
}

func (c *JClass) SetMajorVersion(v uint16) {
	c.majorVersion = v
}
