package jclass

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"strings"
)

// 方法描述符
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

// descriptor => (...parameteTypes)...returnType
// for instance: void fun(int) -> (I)V or ; public int method3(Object obj) -> (Ljava/lang/Object)I
func ParseMethodDescriptor(desc string) *MethodDescriptor {
	methodDesc := MethodDescriptor{raw: desc}
	return methodDesc.parse()
}

func (md *MethodDescriptor) parse() *MethodDescriptor {
	// parse (
	exception.AssertTrue(md.readUint8() == '(', "parse method descriptor error")
	// parse params
	md.parseParamTypes()
	// parse )
	exception.AssertTrue(md.readUint8() == ')', "parse method descriptor error")
	// parse return type
}

//
func (md *MethodDescriptor) parseParamTypes() {
	for {
		t := md.parseFieldType()
		if t != "" {
			md.paramterTypes
		}
	}
}

func (md *MethodDescriptor) addParameterTypes(t string) {

}

func (md *MethodDescriptor) ParamteCount() uint {
	if md.parameteArr == nil {
		md.parameteArr = []string{}
		for range md.parameteTypes {
			md.parameteArr = append(md.parameteArr, md.parseFieldType())
		}
	}

	return uint(len(md.parameteArr))
}

func (md *MethodDescriptor) ParamteTypes() []string {
	if md.parameteArr == nil {
		md.parameteArr = []string{}
		for range md.parameteTypes {
			md.parameteArr = append(md.parameteArr, md.parseFieldType())
		}
	}

	return md.parameteArr
}

func (md *MethodDescriptor) parseFieldType() string {
	switch md.readUint8() {
	case 'B':
		md.offset++
		return "B"
	case 'C':
		md.offset++
		return "C"
	case 'D':
		md.offset++
		return "D"
	case 'F':
		md.offset++
		return "F"
	case 'I':
		md.offset++
		return "I"
	case 'J':
		md.offset++
		return "J"
	case 'S':
		md.offset++
		return "S"
	case 'Z':
		md.offset++
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

func (md *MethodDescriptor) parseObjectType() string {
	unread := md.parameteTypes[md.offset:]
	semicolonIndex := strings.IndexRune(unread, ';')
	exception.AssertTrue(semicolonIndex != -1, "parsing descriptor error")

	objStart := md.offset
	objEnd := md.offset + semicolonIndex
	md.offset = objEnd
	t := md.parameteTypes[objStart:objEnd]
	return t
}

func (md *MethodDescriptor) parseArrayType() string {
	arrStart := md.offset - 1
	md.parseFieldType()
	arrEnd := md.offset
	t := md.parameteTypes[arrStart:arrEnd]
	return t
}

func (md *MethodDescriptor) readUint8() uint8 {
	b := md.paramterTypes[md.offset]
	md.offset++
	return b
}
func (md *MethodDescriptor) unreadUint8() {
	md.offset--
}
