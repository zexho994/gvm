package test

import (
	"github.com/zouzhihao-994/gvm/launcher"
	"testing"
)

// print static field
func TestPrintStaticFields(t *testing.T) {
	launcher.StartGvmByDebug("PrintStaticTest", "/Library/Java/JavaVirtualMachines/jdk1.8.0_261.jdk/Contents/Home/jre", "/Users/zexho/Github/gvm/java/src/test")
}

// print instance field
func TestPrintFields(t *testing.T) {
	launcher.StartGvmByDebug("PrintFieldsTest", "/Library/Java/JavaVirtualMachines/jdk1.8.0_261.jdk/Contents/Home/jre", "/Users/zexho/Github/gvm/java/src/test")
}
