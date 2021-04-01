package constant_pool

import "github.com/zouzhihao-994/gvm/classfile"

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
// 用于表示 references.INVOKE_DYNAMIC 指令用到的以下数据：
// 1. 引导方法
// 2. 引导方法用到的动态调用名称
// 3. 参数
// 4. 返回类型
// 最后还可以为引导方法传入一系列成为静态参数的常量
type ConstantInvokeDynamic struct {
	Tag uint8
	// 对当前class文件中引导方法表 attribute.BootstrapmethodsAttribute 的数组索引
	BootstrapMethodAttrIndex uint16
	// 常量池索引， ConstantNameAndTypeInfo 类型
	NameAndTypeIndex uint16
}

func (invoke *ConstantInvokeDynamic) ReadInfo(reader *classfile.ClassReader) {
	invoke.BootstrapMethodAttrIndex = reader.ReadUint16()
	invoke.NameAndTypeIndex = reader.ReadUint16()
}
