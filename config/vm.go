package config

var JrePath string
var ClassPath string
var ClassName string

var LibJarDir string
var ExtJarDir string

type Options struct {
	MainModule      string
	MainClass       string
	ClassPath       string
	ModulePath      string
	VerboseClass    bool
	VerboseModule   bool
	VerboseJNI      bool
	Xss             string
	Xjre            string
	XUseJavaHome    bool
	XDebugInstr     bool
	XCPUProfile     string
	AbsJavaHome     string // /path/to/jre
	AbsJreLib       string // /path/to/jre/lib
	ThreadStackSize int
}
