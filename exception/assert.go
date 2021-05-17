package exception

// AssertTrue the flag must be true
func AssertTrue(flag bool, msg string) {
	if !flag {
		panic(GvmError{Msg: msg})
	}
}

// AssertFalse the flag must be false
func AssertFalse(flag bool, msg string) {
	if flag {
		panic(GvmError{Msg: msg})
	}
}
