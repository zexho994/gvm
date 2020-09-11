package heap

import "../../classfile"

/*
实例字段
*/
type Field struct {
	ClassMember
	// 字段编号
	slotId uint
}

/*

 */
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfFields := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfFields)
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
