package heap

import "github.com/zouzhihao-994/gvm/src/vm/classfile"

/*
接口方法的符号引用
*/
type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func (self InterfaceMethodRef) Name() string {
	return self.name
}

func (self InterfaceMethodRef) Descriptor() string {
	return self.descriptor
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

/*
解析接口方法
*/
func (self *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if self.method == nil {
		self.resolveInterfaceMethodRef()
	}
	return self.method
}

/*
解析接口方法符号引用
*/
func (self *InterfaceMethodRef) resolveInterfaceMethodRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupInterfaceMethod(c, self.name, self.descriptor)
	if method == nil {

		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}

/*
递归解析方法以及父接口
*/
func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return lookupMethodInInterfaces(iface.interfaces, name, descriptor)
}
