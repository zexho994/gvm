package classfile

import "github.com/zouzhihao-994/gvm/src/share/jclass"

// 类加载阶段的链接阶段
// 链接按照功能细分又可以分为: 验证 -> 准备 -> 解析
func Linked(class *jclass.JClass) {
	verification()
	preparation()
	parse()
}
