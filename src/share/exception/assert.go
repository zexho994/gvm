package exception

// the flag must be true
func AssertTrue(flag bool, msg string) {
	if !flag {
		panic(GvmError{Msg: msg})
	}
}

// the flag must be false
func AssertFalse(flag bool, msg string) {
	if flag {
		panic(GvmError{Msg: msg})
	}
}
