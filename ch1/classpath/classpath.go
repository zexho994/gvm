package classpath

import filepath "path/filepath"
import "os"

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	/*
		解析启动类和扩展类
	*/
	cp.parseBootAndExtClasspath(jreOption)

	/*
		解析用户类t
	*/
	cp.parseUserClasspath(cpOption)
	return cp
}

/*
解析启动类和用户类
启动类 jre/lib/*
扩张类 jre/lib/ext/*
*/
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

/*
查找jre目录的路径
*/
func getJreDir(jreOption string) string {
	// 如果jreOption目录存在
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// 如果 './jre' 目录存在
	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	// 3种情况都不存在jre目录,输出错误
	panic("Can't find jre folder")
}

/*
判断path目录是否存在
*/
func exists(path string) bool {
	// stat获取项目中名称为name的文件信息
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/*
在classpath 中查找 Class文件
*/
func (self Classpath) ReadClass(classpath string) ([]byte, Entry, error) {
	className := classpath + ".class"
	/*

	 */
	if data, entry, err := self.bootClasspath.readClass(className); err != nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err != nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)

}

/*
toString()
打印用户类得路径
*/
func (self Classpath) String() string {
	return self.userClasspath.String()
}

/*
解析用户类
*/
func (self *Classpath) parseUserClasspath(cption string) {
	if cption == "" {
		cption = "."
	}
	self.userClasspath = newEntry(cption)

}
