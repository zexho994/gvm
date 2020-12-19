package exception

func AssertTrue(flag bool, msg string) {
	if !flag {
		panic(GvmError{Msg: msg})
	}
}
