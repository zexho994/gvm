package jclass

import "fmt"

const (
	ACC_PUBLIC       = 0x1    // public class ， 可以从包外访问 ,field method
	ACC_PRIVATE      = 0x2    // field method
	ACC_PROTECTED    = 0x4    // field method
	ACC_STATIC       = 0x8    // field method
	ACC_FINAL        = 0x10   // class field method
	ACC_SUPER        = 0x20   // class
	ACC_SYNCHRONIZED = 0x20   // method
	ACC_VOLATILE     = 0x40   // field
	ACC_BRIDGE       = 0x40   // method
	ACC_TRANSIENT    = 0x80   // field
	ACC_VARARGS      = 0x80   // method
	ACC_NATIVE       = 0x100  // method
	ACC_INTERFACE    = 0x200  // interface class
	ACC_ABSTRACT     = 0x400  // abstract class
	ACC_STRICT       = 0x800  // method
	ACC_SYNTHETIC    = 0x1000 // synthetic，表示class文件并非由java源代码所生成,class field method
	ACC_ANNOTATION   = 0x2000 // 表示注解类型, class
	ACC_ENUM         = 0x4000 // 表示枚举类型, class fiedl
)

func getAccessFlags(acc uint16) {

}

func AccPrint(acc uint16) {
	if IsPublic(acc) {
		fmt.Println("[gvm] class access Is public")
	}
	if IsFinal(acc) {
		fmt.Println("[gvm] class access Is final")
	}
	if IsSuper(acc) {
		fmt.Println("[gvm] class access Is super")
	}
	if IsAbstract(acc) {
		fmt.Println("[gvm] class access Is abstract")
	}
	if IsAnnotation(acc) {
		fmt.Println("[gvm] class access Is annotation")
	}
	if IsEnum(acc) {
		fmt.Println("[gvm] class access Is enum")
	}
	if IsSynthetic(acc) {
		fmt.Println("[gvm] class access Is synthetic")
	}
	if IsInterface(acc) {
		fmt.Println("[gvm] class access Is interface")
	}
	if IsStatic(acc) {
		fmt.Println("[gvm] access flag Is static")
	}
	if IsNative(acc) {
		fmt.Println("[gvm] method access Is native")
	}
}

func IsPublic(acc uint16) bool {
	return (acc & ACC_PUBLIC) != 0
}

func IsFinal(acc uint16) bool {
	return (acc & ACC_FINAL) != 0
}

func IsSuper(acc uint16) bool {
	return (acc & ACC_SUPER) != 0
}

func IsEnum(acc uint16) bool {
	return (acc & ACC_ENUM) != 0
}

func IsInterface(acc uint16) bool {
	return (acc & ACC_INTERFACE) != 0
}

func IsAbstract(acc uint16) bool {
	return (acc & ACC_ABSTRACT) != 0
}

func IsProteced(acc uint16) bool {
	return (acc & ACC_PROTECTED) != 0
}

func IsSynthetic(acc uint16) bool {
	return (acc & ACC_SYNTHETIC) != 0
}

func IsAnnotation(acc uint16) bool {
	return (acc & ACC_ANNOTATION) != 0
}

func IsStatic(acc uint16) bool {
	return (acc & ACC_STATIC) != 0
}

func IsNative(acc uint16) bool {
	return (acc & ACC_NATIVE) != 0
}
