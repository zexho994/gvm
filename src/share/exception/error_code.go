package exception

const (
	// 在实例对象中没有不存在目标字段
	FieldsNotFoundError = "FielsdNotFoundError"
	// 属性表不存在
	AttributeNotFoundError = "AttributeNotFoundError"
	// 解析实例字段过程出现静态字段
	IncompatibleClassChangeError = "IncompatibleClassChangeError"
	// 空指针异常
	NullPointException = "NullPointException"
)
