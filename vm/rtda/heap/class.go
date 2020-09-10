package heap

import "../../classfile"

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
	//
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        *Slots
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
