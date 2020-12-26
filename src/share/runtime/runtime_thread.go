package runtime

// 映射到java中的一个thread todo
type Thread struct {
	PC uint
	*Stack
}
