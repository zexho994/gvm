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
	if md.parameteArr != nil {
		return uint(len(md.parameteArr))
	}
	md.parameteArr = []string{}
	for range md.parameteTypes {
		md.parameteArr = append(md.parameteArr, md.parseFieldType())
	}

	return uint(len(md.parameteArr))
}

func (self *MethodDescriptor) parseFieldType() string {
	switch self.parameteTypes[self.offset] {
	case 'B':
		self.offset++
		return "B"
	case 'C':
		self.offset++
		return "C"
	case 'D':
		self.offset++
		return "D"
	case 'F':
		self.offset++
		return "F"
	case 'I':
		self.offset++
		return "I"
	case 'J':
		self.offset++
		return "J"
	case 'S':
		self.offset++
		return "S"
	case 'Z':
		self.offset++
		return "Z"
	case 'L':
		return self.parseObjectType()
	case '[':
		return self.parseArrayType()
	default:
		return ""
	}
}

func (self *MethodDescriptor) parseObjectType() string {
	unread := self.parameteTypes[self.offset:]
	semicolonIndex := strings.IndexRune(unread, ';')
	exception.AssertTrue(semicolonIndex == -1, "parsing descriptor error")

	objStart := self.offset
	objEnd := self.offset + semicolonIndex + 1
	self.offset = objEnd
	t := self.parameteTypes[objStart:objEnd]
	return t
}

func (self *MethodDescriptor) parseArrayType() string {
	arrStart := self.offset
	self.parseFieldType()
	arrEnd := self.offset
	t := self.parameteTypes[arrStart:arrEnd]
	return t
}
