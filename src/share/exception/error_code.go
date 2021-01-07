package exception

const (
	// 在实例对象中没有不存在目标字段
	FIELDS_NOT_FOUND_ERROR = "FielsdNotFound"
	// 解析实例字段过程出现静态字段
	INCOMPATIBLE_CLASS_CHANGE_ERROR = "IncompatibleClassChangeError"
	// 空指针异常
	NULL_POINT_EXCEPTION = "NullPointException"
)
