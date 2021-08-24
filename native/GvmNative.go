package native

import "strings"

func isGvmNative(methodName string) bool {
	if strings.Contains(methodName, "GvmOut") {
		return true
	}

	return false
}
