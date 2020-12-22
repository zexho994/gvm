package jclass

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"strings"
)

type MethodDescriptor struct {
	parameteTypes string
	parameteArr   []string
	returnType    string
	offset        int
}

// descriptor => (...parameteTypes)...returnType
// for instance: void fun(int) -> (I)V or ; public int method3(Object obj) -> (Ljava/lang/Object)I
func ParseMethodDescriptor(method *MethodInfo) *MethodDescriptor {
	idx := method.descriptorIdx
	descStr := method.CP().GetUtf8(idx)
	splits := strings.Split(descStr, ")")
	parametes := strings.Split(splits[0], "(")[0]

	return &MethodDescriptor{
		parameteTypes: parametes,
		returnType:    splits[1],
	}
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
	switch md.parameteTypes[md.offset] {
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
		return ""
	}
}

func (md *MethodDescriptor) parseObjectType() string {
	unread := md.parameteTypes[md.offset:]
	semicolonIndex := strings.IndexRune(unread, ';')
	exception.AssertTrue(semicolonIndex == -1, "parsing descriptor error")

	objStart := md.offset
	objEnd := md.offset + semicolonIndex + 1
	md.offset = objEnd
	t := md.parameteTypes[objStart:objEnd]
	return t
}

func (md *MethodDescriptor) parseArrayType() string {
	arrStart := md.offset
	md.parseFieldType()
	arrEnd := md.offset
	t := md.parameteTypes[arrStart:arrEnd]
	return t
}
