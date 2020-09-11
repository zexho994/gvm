package heap

import "../../classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

/**
字段符号引用解析
*/
func (self *FieldRef) ResolveField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

/*
解析字段符号引用
*/
func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	field := lookupField(c, self.name, self.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.field = field
}

/**
循环查找字段
在类c中根据描述父查询
查询逻辑 => 查询本类 -> 查询接口类 -> 查询父类
*/
func lookupField(c *Class, name, descriptor string) *Field {
	// 遍历字段表
	for _, field := range c.fields {
		// 名称相同而且描述符一样
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	// 遍历接口
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}
	if c.superClass != nil {
		// 在父类中查询
		return lookupField(c.superClass, name, descriptor)
	}
	return nil
}

/**
字段访问规则
public可以
protected同一子类或者同一包都可以
private需要同一个包下
*/
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}
