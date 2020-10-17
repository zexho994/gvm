package heap

import "github.com/zouzhihao-994/gvm/src/vm/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self MethodRef) Name() string {
	return self.name
}

func (self MethodRef) Descriptor() string {
	return self.descriptor
}

/*
解析非接口方法
*/
func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

/*
解析非接口方法符号引用
*/
func (self *MethodRef) resolveMethodRef() {
	// 获取方法的类
	d := self.cp.class
	// 解析类
	c := self.ResolvedClass()
	// 如果类是接口，则不能其调用方法
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 在目标类中遍历搜寻该方法
	method := lookupMethod(c, self.name, self.descriptor)
	// 如果找不到对应的方法
	if method == nil {

		panic("java.lang.NoSuchMethodError")
	}
	// 如果不能访问对应的类
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	// 设置方法
	self.method = method
}

/*
循环搜索指定的方法
*/
func lookupMethod(class *Class, name, descriptor string) *Method {
	// 先在类和父类中搜寻
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		// 类和父类中搜寻不到在接口中搜寻
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}

/*
在类和父类中搜寻指定的方法
*/
func LookupMethodInClass(class *Class, name, descriptor string) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			// 判断类相同的条件：1 名称相同 2描述符相同
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

/*
在接口中搜寻
*/
func lookupMethodInInterfaces(ifaces []*Class, name, descriptor string) *Method {
	// 遍历接口
	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method := lookupMethodInInterfaces(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}
