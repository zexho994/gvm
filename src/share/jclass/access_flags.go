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
	if isPublic(acc) {
		fmt.Println("[gvm] class access is public")
	}
	if isFinal(acc) {
		fmt.Println("[gvm] class access is final")
	}
	if isSuper(acc) {
		fmt.Println("[gvm] class access is super")
	}
	if isAbstract(acc) {
		fmt.Println("[gvm] class access is abstract")
	}
	if isAnnotation(acc) {
		fmt.Println("[gvm] class access is annotation")
	}
	if isEnum(acc) {
		fmt.Println("[gvm] class access is enum")
	}
	if isSynthetic(acc) {
		fmt.Println("[gvm] class access is synthetic")
	}
	if isInterface(acc) {
		fmt.Println("[gvm] class access is interface")
	}
}

func isPublic(acc uint16) bool {
	return (acc & ACC_PUBLIC) != 0
}

func isFinal(acc uint16) bool {
	return (acc & ACC_FINAL) != 0
}

func isSuper(acc uint16) bool {
	return (acc & ACC_SUPER) != 0
}

func isEnum(acc uint16) bool {
	return (acc & ACC_ENUM) != 0
}

func isInterface(acc uint16) bool {
	return (acc & ACC_INTERFACE) != 0
}

func isAbstract(acc uint16) bool {
	return (acc & ACC_ABSTRACT) != 0
}

func isSynthetic(acc uint16) bool {
	return (acc & ACC_SYNTHETIC) != 0
}

func isAnnotation(acc uint16) bool {
	return (acc & ACC_ANNOTATION) != 0
}

func isStatic(acc uint16) bool {
	return (acc & ACC_STATIC) != 0
}
