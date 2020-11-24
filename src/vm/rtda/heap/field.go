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

func (field Field) Class() *Class {
	return field.class
}

func (field Field) ConstValueIndex() uint {
	return field.constValueIndex
}

func (field Field) IsStatic() bool {
	return 0 != field.access&ACC_STATIC
}

func (field Field) Descriptor() string {
	return field.descriptor
}

func (field Field) SlotId() uint {
	return field.slotId
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
func (field Field) IsLongOrDouble() bool {
	return field.descriptor == "J" || field.descriptor == "D"
}

func (field Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		field.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (field Field) IsFinal() bool {
	return 0 != field.access&ACC_FINAL
}
