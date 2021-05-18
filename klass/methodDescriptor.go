package klass

import (
	"github.com/zouzhihao-994/gvm/utils"
	"strings"
)

// MethodDescriptor 方法描述符
type MethodDescriptor struct {
	// 描述符原始字段
	raw string
	// 偏移量
	offset int
	// 参数数组
	paramterTypes []string
	// 返回类型
	returnTypt string
}

func (md *MethodDescriptor) Paramters() []string {
	return md.paramterTypes
}

// ParseMethodDescriptor descriptor => (...parameteTypes)...returnType
// for instance: void fun(int) -> (I)V or ; public int method3(Object obj) -> (Ljava/lang/Object)I
func ParseMethodDescriptor(desc string) *MethodDescriptor {
	methodDesc := MethodDescriptor{raw: desc}
	return methodDesc.parse()
}

func (md *MethodDescriptor) parse() *MethodDescriptor {
	// parse (
	utils.AssertTrue(md.readUint8() == '(', "parse method descriptor error")
	// parse params
	md.parseParamTypes()
	// parse )
	utils.AssertTrue(md.readUint8() == ')', "parse method descriptor error")
	// parse return type
	md.parseReturnType()
	utils.AssertTrue(md.offset == len(md.raw), "parse mthod descriptor error")
	return md
}

//
func (md *MethodDescriptor) parseParamTypes() {
	for {
		t := md.parseFieldType()
		if t != "" {
			md.addParameterTypes(t)
		} else {
			break
		}
	}
}

func (md *MethodDescriptor) ParamsCount() uint {
	return uint(len(md.paramterTypes))
}

func (md *MethodDescriptor) addParameterTypes(t string) {
	pLen := len(md.paramterTypes)
	if pLen == cap(md.paramterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, md.paramterTypes)
		md.paramterTypes = s
	}

	md.paramterTypes = append(md.paramterTypes, t)
}

func (md *MethodDescriptor) parseFieldType() string {
	switch md.readUint8() {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return md.parseObjectType()
	case '[':
		return md.parseArrayType()
	default:
		md.unreadUint8()
		return ""
	}
}

func (md *MethodDescriptor) parseReturnType() {
	if md.readUint8() == 'V' {
		md.returnTypt = "V"
		return
	}
	md.unreadUint8()
	t := md.parseFieldType()
	utils.AssertTrue(t != "", "parse return type error")
	md.returnTypt = t
}

func (md *MethodDescriptor) parseObjectType() string {
	unread := md.raw[md.offset:]
	semicolonIndex := strings.IndexRune(unread, ';')
	utils.AssertFalse(semicolonIndex == -1, "parsing descriptor error")

	objStart := md.offset - 1
	objEnd := md.offset + semicolonIndex + 1
	md.offset = objEnd
	t := md.raw[objStart:objEnd]
	return t
}

func (md *MethodDescriptor) parseArrayType() string {
	arrStart := md.offset - 1
	md.parseFieldType()
	arrEnd := md.offset
	t := md.raw[arrStart:arrEnd]
	return t
}

func (md *MethodDescriptor) readUint8() uint8 {
	b := md.raw[md.offset]
	md.offset++
	return b
}
func (md *MethodDescriptor) unreadUint8() {
	md.offset--
}
