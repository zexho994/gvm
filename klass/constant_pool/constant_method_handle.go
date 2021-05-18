package constant_pool

import (
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/loader"
)

// ConstantMethodHandleInfo
// 是一个强类型，可以直接指向的引用
// 可以指向静态方法、实例方法、构造器或者字段
// 指向字段时：实则指向包含字段访问的虚方法，语义上等价于目标自段的getter/setter方法
// 指向方法时：
//
// CONSTANT_MethodHandle_info {
//     u1 tag;
//     u1 reference_kind;
//     u2 reference_index;
// }
type ConstantMethodHandleInfo struct {
	Tag uint8
	// 方法句柄类型,句柄类型决定了字节码j行为
	// 1.ref_getField,2.ref_getStatic,3.ref_putField
	// 4.ref_putStatic,5.ref_invokeVirtual,6.ref_invokeStatic
	// 7.ref_invokeSpeical,8.ref_newInvokeSpecial,9.ref_infokeInterface
	ReferenceKind uint8
	// 常量池索引
	// 对于 ReferenceKind 值,必须满足一下条件
	// - 如果为1，2，3，4，那么索引对应的类型必须是 ConstantFieldInfo,表示某个字段，本方法句柄正是为这个字段创建
	// - 如果为5,8,那么索引类型为 ConstantMethodInfo,表示某个方法或者构造器
	// - 如果为6、7且版本号小于52.0,索引对应类型为 ConstantMethodInfo 或者 ConstantInterfaceMethodInfo
	// - 如果为9，那么类型必须为 ConstantInterfaceMethodInfo
	// - 如果为5，6，7，9，那么必须为 ConstantMethodInfo 或者 ConstantInterfaceMethodInfo,名称不能为<init>,<clinit>
	// - 如果为8，那么必须为 ConstantMethodInfo 结构表示的方法 ，名称必须是<init>
	ReferenceIndex uint16
	cp             ConstantPool
}

func (handle *ConstantMethodHandleInfo) ReadInfo(reader *loader.ClassReader) {
	handle.ReferenceKind = reader.ReadUint8()
	handle.ReferenceIndex = reader.ReadUint16()
}

// 解析出 ReferenceKind 对应的字节码行为.
// x，T分别表示字段或方法的名称和描述符，C表示字段或方法所属的类或者接口
// * ref_getField -> getfield C.f:T
// * ref_getStatic -> getstatic C.f:T
// * ref_putField -> putfield C.f:T
// * ref_putStatic -> putstatic C.f:T
// * ref_invokeVirtual -> invokevirtual C.m:(A*)T
// * ref_invokeStatic -> invokestatic C.m:(A*)T
// * ref_invokeSPecial -> invokespeical C.m:(A*)T
// * ref_newInvokeSpecial -> new C;dup;invokespecial C.<init>:(A*)void
// * ref_invokeInterface -> invokeinterface C.m:(A*)T
func (handle ConstantMethodHandleInfo) ParseKindRef() {
	// 1. 解析R，R为handle中字段或者方法的符号引用
	ref := handle.cp.GetConstantInfo(handle.ReferenceIndex)
	if handle.ReferenceKind == 6 {
		methodRef := ref.(*ConstantMethodInfo)
		if name, _ := methodRef.NameAndDescriptor(); name == "init" || name == "clinit" {
			exception.GvmError{Msg: exception.MethodParseException}.Throw()
		}
	}

	// 2. 按照解析指向类和接口的未解析符号引用的步骤来解析这些符号引用
	//    类和接口的名称分别对应于A*中每个类型以及类型T
	//    解析的顺序也是先解析A*的符号引用，再解析指向T的符号引用

	// 3. 像解析指向方法类型的为解析符号引用时那样，来解析并获取指向java.lang.invoke.MethodType实例的引用

}
