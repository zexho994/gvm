package heap

/*
符号引用
cp,className,class是所有符号引用通用的字段
利用继承的方式来进行定制
*/
type SymRef struct {
	// 常量池指针
	cp *ConstantPool
	// 类完全限定名
	className string
	// 解析后的类结构体指针
	class *Class
}
