package classfile

// 类加载阶段的链接阶段
// 链接按照功能细分又可以分为: 验证 -> 准备 -> 解析
func Linked() {
	verification()
	preparation()
	parse()
}
