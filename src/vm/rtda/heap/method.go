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

func (method *Method) Class() *Class {
	return method.class
}

func (method *Method) Code() []byte {
	return method.code
}

func (method Method) MaxStack() uint {
	return method.maxStack
}

func (method Method) MaxLocals() uint {
	return method.maxLocals
}

func (method Method) IsStatic() bool {
	return 0 != method.access&ACC_STATIC
}

func (method Method) Name() string {
	return method.name
}

func (method Method) MethodDescriptor() string {
	return method.descriptor
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

func (method *Method) injectCodeAttribute(returnType string) {
	method.maxStack = 4 // todo
	method.maxLocals = method.argSlotCount
	switch returnType[0] {
	case 'V':
		method.code = []byte{0xfe, 0xb1} // return
	case 'L', '[':
		method.code = []byte{0xfe, 0xb0} // areturn
	case 'D':
		method.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		method.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		method.code = []byte{0xfe, 0xad} // lreturn
	default:
		method.code = []byte{0xfe, 0xac} // ireturn
	}
}

func (method *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		method.maxLocals = codeAttr.MaxLocals()
		method.maxStack = codeAttr.MaxStack()
		method.code = codeAttr.Code()
	}
}

func (method *Method) IsNative() bool {
	return 0 != method.access&ACC_NATIVE
}

func (method *Method) ArgSlotCount() uint { return method.argSlotCount }

/*
计算参数数量
*/
func (method *Method) calcArgSlotCount(paramTypes []string) {
	// 解析方法的描述符
	for _, paramType := range paramTypes {
		method.argSlotCount++
		// long和double类型要额外1个空间
		if paramType == "J" || paramType == "D" {
			method.argSlotCount++
		}
	}

	if !method.IsStatic() {
		method.argSlotCount++
	}
}

func (method Method) IsAbstract() bool {
	return 0 != method.access&ACC_ABSTRACT
}

func (method Method) Descriptor() string {
	return method.descriptor
}
