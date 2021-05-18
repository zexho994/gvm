package utils

import "github.com/zouzhihao-994/gvm/exception"

// AssertTrue the flag must be true
func AssertTrue(flag bool, msg string) {
	if !flag {
		panic(exception.GvmError{Msg: msg})
	}
}

// AssertFalse the flag must be false
func AssertFalse(flag bool, msg string) {
	if flag {
		panic(exception.GvmError{Msg: msg})
	}
}
