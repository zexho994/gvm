package jclass

import "fmt"

const (
	ACC_PUBLIC      = 0x1  // public class ， 可以从包外访问
	ACC_FINAL       = 0x10 // final class , 不许有子类
	ACC_SUPER       = 0x20
	ACC_INTERFACE   = 0x200  // interface class
	ACC_ABSTRACT    = 0x400  // abstract class
	ACC_SYNTHETIC   = 0x1000 // synthetic，表示class文件并非由java源代码所生成
	ACC_ACNNOTATION = 0x2000 // 表示注解类型
	ACC_ENUM        = 0x4000 // 表示枚举类型
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
