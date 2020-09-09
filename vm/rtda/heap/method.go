package heap

import "../../classfile"

type Method struct {
	ClassMember
	maxStack  uint16
	maxLocals uint16
	// 存放方法表的Code字段
	code []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxLocals = codeAttr.MaxLocals()
		self.maxStack = codeAttr.MaxStack()
		self.code = codeAttr.Code()
	}
}
