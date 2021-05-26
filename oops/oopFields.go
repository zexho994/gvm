package oops

import (
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/utils"
)

// OopFields 实例对象中的实例字段表
type OopFields []OopField

type OopField struct {
	// 实例对象字段名称
	name       string
	accessFlag uint16
	slots      utils.Slots
}

func FindField(name string, fields *OopFields, k *klass.Klass) *OopField {
	if k == nil {
		return nil
	}
	f, r := fields.GetField(name)
	if r {
		return &f
	}
	fields = InitOopFields(k)
	return FindField(name, fields, k.SuperClass)
}

// GetField 查找实例字段
// 如果本类中找不到，就在父类中找
// name:字段名称
// isSuper：是否是从子类中进行调用的
func (fields OopFields) GetField(name string) (OopField, bool) {
	for idx := range fields {
		if fields[idx].name != name {
			continue
		}
		if utils.IsFinal(fields[idx].accessFlag) {
			// todo
			//exception.GvmError{Msg: "final fields not be inheritance"}.Throw()
		}
		return fields[idx], true
	}
	return OopField{}, false
}

// InitOopFields 初始化实例对象的实例字段表
func InitOopFields(instance *klass.Klass) *OopFields {
	fields := &OopFields{}
	jf := instance.Fields
	for idx := 0; idx < len(jf); idx++ {
		flag := jf[idx].AccessFlags
		if utils.IsStatic(flag) {
			continue
		}
		desc := jf[idx].Descriptor()
		slots := utils.Slots{}
		slot := utils.Slot{Type: utils.TypeMapping(desc)}
		if desc == "D" || desc == "J" {
			slots = append(slots, slot)
		}
		slots = append(slots, slot)
		newField := OopField{name: jf[idx].Name(), accessFlag: flag, slots: slots}
		*fields = append(*fields, newField)
	}

	return fields
}

func (field OopField) Name() string {
	return field.name
}

func (field OopField) Slots() utils.Slots {
	return field.slots
}

func (field OopField) AccessFlag() uint16 {
	return field.accessFlag
}
