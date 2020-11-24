package loader

import (
	"path/filepath"
)
import "os"

// 加载器
type Loader struct {
	// 启动类加载器
	bootLoader Entry
	// 扩张类加载器
	extLoader Entry
	// 应用类加载器
	userLoader Entry
}

// 将-Xjre 和 class 两个字段进行解析
// xJre : 启动类和扩展类路径
// cp/loader : 用户类路径
func Parse(jreOption, cpOption string) *Loader {
	// 创建一个新的Classpath类返回其地址
	loader := &Loader{}

	// 解析启动类和扩展类
	// fmt.Printf("[gvm][Parse] jreOption : %v\n", jreOption)
	loader.parseBootAndExtLoader(jreOption)

	// 解析用户类
	// fmt.Printf("[gvm][Parse] cpOtion : %v\n", cpOption)
	loader.parseUserLoader(cpOption)
	return loader
}

// 解析启动类和用户类
// 启动类 {jre}/lib/*
// 扩张类 {jre}/lib/ext/*
func (c *Loader) parseBootAndExtLoader(jreOption string) {
	// 查找jre目录路径
	jreDir := getJreDir(jreOption)
	//fmt.Printf("[gvm][parseBootAndExtLoader] jreDir : %v\n", jreDir)

	// 拼接/lib/*目录 , 然后创建wildcardEntry 对象
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	//fmt.Printf("[gvm][parseBootAndExtLoader] jreLibDir : %v\n", jreLibPath)
	// 设置应用类加载器
	c.bootLoader = newWildcardEntry(jreLibPath)

	// 拼接lib/ext/* 目录 , 然后创建wildcardEntry 对象
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	//fmt.Printf("[gvm][parseBootAndExtLoader] jreExtDir : %v\n", jreExtPath)
	// 设置扩展类加载器
	c.extLoader = newWildcardEntry(jreExtPath)
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

// 查找用户类目录
func (c *Loader) parseUserLoader(cption string) {
	// 如果没有用户没有输入 -cp ,默认当前目录为用户目录
	if cption == "" {
		cption = "."
	}
	// 设置应用类加载器
	c.userLoader = newEntry(cption)
}

/*
在classpath 中查找 Class文件
*/
func (c Loader) ReadClass(classpath string) ([]byte, Entry, error) {
	// 拼接类名
	className := classpath + ".class"

	// 从className中读取 bootLoader
	if data, entry, err := c.bootLoader.readClass(className); err == nil {
		//fmt.Printf("[gvm][ReadClss] return bootLoader <data> : %v\n", data)
		return data, entry, err
	}

	// 从className中读取 extLoader
	if data, entry, err := c.extLoader.readClass(className); err == nil {
		//fmt.Printf("[gvm][ReadClss] return extLoader <data> : %v\n", data)
		return data, entry, err
	}

	// 从className中读取 userLoader
	return c.userLoader.readClass(className)

}

/*
toString()
打印用户类得路径
*/
func (c Loader) String() string {
	return c.userLoader.String()
}