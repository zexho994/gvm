package heap

import (
	"../../classfile"
	"strings"
)

type Class struct {
	// 类访问标志，public,protect,private
	accessFlags uint16
	// 类完全限定名
	name string
	// 父类完全限定名
	superClassName string
	// 接口完全限定名
	interfaceNames []string
	// 常量池指针
	constantPool *ConstantPool
	// 字段表
	fields []*Field
	// 方法表
	methods []*Method
	// 类加载器
	loader *ClassLoader
	// 父类
	superClass *Class
	// 接口
	interfaces []*Class
	// 实例字段数量
	instanceSlotCount uint
	// 静态字段数量
	staticSlotCount uint
	staticVars      *Slots
}

/*
将字节码信息转化成Class结构
*/
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *Class) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}

func (self *Class) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}

/**
一个类想要访问另一个类，必须含有
1. 两个类在同一个包下
2. 要访问的类public
*/
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.getPackageName() == other.getPackageName()

}

/**
获取包名
lastIndex : "/" 在 name中出现的最后所以下标
如果name = "java/lang/Object","/"出现的最后位置就是9
*/
func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}
