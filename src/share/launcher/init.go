package launcher

// 在这里决定使用那种方法启动gvm
func InitializeGVM() {
	StartGvmByDebug("Ldc", "/Library/Java/JavaVirtualMachines/jdk1.8.0_261.jdk/Contents/Home/jre", "java/src")
}
