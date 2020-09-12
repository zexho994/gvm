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

/**
解析符号引用
*/
func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		// 解析类符号引用
		self.resolveClassRef()
	}
	// 如果类已经被解析类，直接返回类指针
	return self.class
}

func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}
