package methodArea

import "github.com/zouzhihao-994/gvm/src/vm/classfile"

type Perm struct {
	classMap map[string]*classfile.ClassFile
}

func (perm Perm) FindClass(className string) *classfile.ClassFile {
	if class, ok := perm.classMap[className]; ok {
		return class
	}
	return nil
}
