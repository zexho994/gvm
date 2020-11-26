package oops

import "github.com/zouzhihao-994/gvm/src/vm/classfile"

type ClassMember struct {
	access     uint16
	name       string
	descriptor string
	//Class结构体指针,这样可以通过字段和方法访问到类
	class *Class
}

/*
复制
*/
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.access = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.access&ACC_PUBLIC
}

func (self *ClassMember) IsPrivate() bool {
	return 0 != self.access&ACC_PRIVATE
}

func (self *ClassMember) IsProtected() bool {
	return 0 != self.access&ACC_PROTECTED
}

// jvms 5.4.4
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.IsSubClassOf(c) ||
			c.GetPackageName() == d.GetPackageName()
	}
	if !self.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}
	return d == c
}
