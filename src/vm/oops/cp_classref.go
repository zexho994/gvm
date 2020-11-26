package oops

import "github.com/zouzhihao-994/gvm/src/vm/classfile"

type ClassRef struct {
	SymRef
}

func newClssRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}
