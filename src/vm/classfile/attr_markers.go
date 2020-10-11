package classfile

/*
@Derecated 注解可以标示方法或者类已经过期
*/
type DeprecatedAttribute struct{ MarkerAttribute }

/*
标记源文件不存在,由编译器生成的类成员
*/
type SyntheticAttribute struct{ MarkerAttribute }
type MarkerAttribute struct{}

/*
由于deprecater和synthetic两者只起到标记的作用,不包含数据.
*/
func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
