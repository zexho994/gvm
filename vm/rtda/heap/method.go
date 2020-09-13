package heap

import "../../classfile"

type Method struct {
	ClassMember
	maxStack  uint16
	maxLocals uint16
	// 存放方法表的Code字段
	code []byte
}

func (self *Method) Class() *Class {
	return self.class
}

func (self *Method) Code() []byte {
	return self.code
}

func (self Method) MaxStack() uint16 {
	return self.maxStack
}

func (self Method) MaxLocals() uint16 {
	return self.maxLocals
}

func (self Method) IsStatic() bool {
	return 0 != self.access&ACC_STATIC
}

func (self Method) Name() string {
	return self.name
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
