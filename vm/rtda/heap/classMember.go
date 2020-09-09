package heap

import "../../classfile"

type ClassMember struct {
	access uint16
	name   string
	/*

	 */
	descriptor string
	/*
		Class结构体指针
		这样可以通过字段和方法访问到类
	*/
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
