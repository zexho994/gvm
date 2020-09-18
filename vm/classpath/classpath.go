package classpath

import (
	"path/filepath"
)
import "os"

type Classpath struct {
	// 启动类加载器
	bootClasspath Entry
	// 扩张类加载器
	extClasspath Entry
	// 应用类加载器
	userClasspath Entry
}

/*
将-Xjre 和 class 两个字段进行解析
xJre : 启动类和扩展类路径
cp/classpath : 用户类路径
*/
func Parse(jreOption, cpOption string) *Classpath {
	// 创建一个新的Classpath类返回其地址
	cp := &Classpath{}
	/*
		解析启动类和扩展类
	*/
	//fmt.Printf("[gvm][Parse] jreOption : %v\n", jreOption)
	cp.parseBootAndExtClasspath(jreOption)

	/*
		解析用户类
	*/
	//fmt.Printf("[gvm][Parse] cpOtion : %v\n", cpOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

/*
解析启动类和用户类
启动类 jre/lib/*
扩张类 jre/lib/ext/*
*/
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	// 查找jre目录路径
	jreDir := getJreDir(jreOption)
	//fmt.Printf("[gvm][parseBootAndExtClasspath] jreDir : %v\n", jreDir)

	// 拼接/lib/*目录 , 然后创建wildcardEntry 对象
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	//fmt.Printf("[gvm][parseBootAndExtClasspath] jreLibDir : %v\n", jreLibPath)
	// 设置应用类加载器
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// 拼接lib/ext/* 目录 , 然后创建wildcardEntry 对象
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	//fmt.Printf("[gvm][parseBootAndExtClasspath] jreExtDir : %v\n", jreExtPath)
	// 设置扩展类加载器
	self.extClasspath = newWildcardEntry(jreExtPath)
}

/*
查找jre目录的路径
*/
func getJreDir(jreOption string) string {
	// 如果用户输入了-Xjre 参数
	if jreOption != "" && exists(jreOption) {
		//fmt.Printf("[gvm][getJreDir] 获取到/jre的路径 : %v\n", jreOption)
		return jreOption
	}
	/*
		如果用户没有输入 -Xjre 参数
		在 './jre' 下找
	*/
	if exists("./jre") {
		//fmt.Printf("[gvm][getJreDir] 获取到/jre的路径\n")
		return "./jre"
	}

	/**
	没有输入 -Xjre参数,而且当前目录下也没有找到
	在 JAVA_HOME中找
	*/
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		//fmt.Printf("[gvm][getJreDir] find JAVA_HOME dir")
		return filepath.Join(jh, "jre")
	}

	// 3种情况都不存在jre目录,输出错误
	panic("[gvm][getJreDir]Can't find jre folder")
}

/*
判断path目录是否存在
*/
func exists(path string) bool {
	// stat获取项目中path的文件信息
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/*
查找用户类目录
*/
func (self *Classpath) parseUserClasspath(cption string) {
	// 如果没有用户没有输入 -cp ,默认当前目录为用户目录
	if cption == "" {
		cption = "."
	}
	// 设置应用类加载器
	self.userClasspath = newEntry(cption)
}

/*
在classpath 中查找 Class文件
*/
func (self Classpath) ReadClass(classpath string) ([]byte, Entry, error) {
	// 拼接类名
	className := classpath + ".class"

	// 从className中读取 bootClasspath
	//fmt.Printf("[gvm][ReadClss] to read %v from bootClasspath\n", className)
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		//fmt.Printf("[gvm][ReadClss] return bootClasspath <data> : %v\n", data)
		return data, entry, err
	}

	// 从className中读取 extClasspath
	//fmt.Printf("[gvm][ReadClss] to read %v from extClasspath\n", className)
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		//fmt.Printf("[gvm][ReadClss] return extClasspath <data> : %v\n", data)
		return data, entry, err
	}

	// 从className中读取 userClasspath
	//fmt.Printf("[gvm][ReadClss] to read %v from userClasspath \n", className)
	return self.userClasspath.readClass(className)

}

/*
toString()
打印用户类得路径
*/
func (self Classpath) String() string {
	return self.userClasspath.String()
}
