package heap

import "github.com/zouzhihao-994/gvm/src/vm/classfile"

/*
实例字段
*/
type Field struct {
	ClassMember
	// 常量索引
	constValueIndex uint
	// 字段编号
	slotId uint
}

func (self Field) Class() *Class {
	return self.class
}

func (self Field) ConstValueIndex() uint {
	return self.constValueIndex
}

func (self Field) IsStatic() bool {
	return 0 != self.access&ACC_STATIC
}

func (self Field) Descriptor() string {
	return self.descriptor
}

func (self Field) SlotId() uint {
	return self.slotId
}

/*

 */
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfFields := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfFields)
		fields[i].copyAttributes(cfFields)
	}
	return fields
}

/*
判断描述符是否属于J或者D，J -> long,D -> double
基本类型中，除了long的定义奇特外，其他都是基于首字母
*/
func (self Field) IsLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self Field) IsFinal() bool {
	return 0 != self.access&ACC_FINAL
}
