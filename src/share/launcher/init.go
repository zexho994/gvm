package launcher

// 在这里决定使用那种方法启动gvm
func InitializeGVM() {
	StartGvmByDebug("CreateArray", "/Library/Java/JavaVirtualMachines/jdk1.8.0_261.jdk/Contents/Home/jre", "java/src/array")
}
