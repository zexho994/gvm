package oops

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/utils"
)

// 实例对象中的实例字段表
type OopFields []OopField

type OopField struct {
	// 实例对象字段名称
	name       string
	accessFlag uint16
	slots      utils.Slots
}

func FindField(name string, fields *OopFields, instance *Oop_Instance, isSuper bool) *OopField {
	f := fields.GetField(name, isSuper)
	if f != nil {
		return f
	}
	fields = InitOopFields(instance.jclassInstance.SuperClass)
	return FindField(name, fields, instance, true)
}

// 查找实例字段
// 如果本类中找不到，就在父类中找
func (fields OopFields) GetField(name string, isSuper bool) *OopField {
	for idx := range fields {
		if fields[idx].name != name {
			continue
		}
		if jclass.IsFinal(fields[idx].accessFlag) && isSuper {
			exception.GvmError{Msg: "final field not be inheritance"}.Throw()
		}
		return &fields[idx]
	}
	return nil
}

// 初始化实例对象的实例字段表
func InitOopFields(instance *jclass.JClass_Instance) *OopFields {
	fields := OopFields{}
	jf := instance.Fields
	for idx := range jf {
		flag := jf[idx].AccessFlags
		if jclass.IsStatic(flag) {
			continue
		}
		name := jf[idx].Name()
		desc := jf[idx].Descriptor()
		slots := utils.Slots{}
		slot := utils.Slot{Type: utils.TypeMapping(desc)}
		// double & long 需要两个slot
		if desc == "D" || desc == "J" {
			slots = append(slots, slot)
		}
		slots = append(slots, slot)
		newField := OopField{name: name, accessFlag: flag, slots: slots}
		fields = append(fields, newField)
	}

	return &fields
}

func (filed OopField) Name() string {
	return filed.name
}

func (field OopField) Slots() utils.Slots {
	return field.slots
}

func (field OopField) AccessFlag() uint16 {
	return field.accessFlag
}
