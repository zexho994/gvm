package heap

import "github.com/zouzhihao-994/gvm/src/vm/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	// 存放方法表的Code字段
	code []byte
	// 参数数量
	argSlotCount uint
}

func (self *Method) Class() *Class {
	return self.class
}

func (self *Method) Code() []byte {
	return self.code
}

func (self Method) MaxStack() uint {
	return self.maxStack
}

func (self Method) MaxLocals() uint {
	return self.maxLocals
}

func (self Method) IsStatic() bool {
	return 0 != self.access&ACC_STATIC
}

func (self Method) Name() string {
	return self.name
}

func (self Method) MethodDescriptor() string {
	return self.descriptor
}

/*
方法表
*/
func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

/*
单个方法
*/
func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

func (self *Method) injectCodeAttribute(returnType string) {
	self.maxStack = 4 // todo
	self.maxLocals = self.argSlotCount
	switch returnType[0] {
	case 'V':
		self.code = []byte{0xfe, 0xb1} // return
	case 'L', '[':
		self.code = []byte{0xfe, 0xb0} // areturn
	case 'D':
		self.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		self.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		self.code = []byte{0xfe, 0xad} // lreturn
	default:
		self.code = []byte{0xfe, 0xac} // ireturn
	}
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxLocals = codeAttr.MaxLocals()
		self.maxStack = codeAttr.MaxStack()
		self.code = codeAttr.Code()
	}
}

func (self *Method) IsNative() bool {
	return 0 != self.access&ACC_NATIVE
}

func (self *Method) ArgSlotCount() uint { return self.argSlotCount }

/*
计算参数数量
*/
func (self *Method) calcArgSlotCount(paramTypes []string) {
	// 解析方法的描述符
	for _, paramType := range paramTypes {
		self.argSlotCount++
		// long和double类型要额外1个空间
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}

	if !self.IsStatic() {
		self.argSlotCount++
	}
}

func (self Method) IsAbstract() bool {
	return 0 != self.access&ACC_ABSTRACT
}

func (self Method) Descriptor() string {
	return self.descriptor
}
