package heap

import "github.com/zouzhihao-994/gvm/src/vm/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

/*
将class文件常量池中方法与字段的信息拷贝到成员结构体中
*/
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}
