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
func New(jreOption, cpOption string) *Loader {
	// 创建一个新的Classpath类返回其地址
	loader := &Loader{}
	// 解析启动类，/lib
	libPath := loader.parseBootLoader(jreOption)
	// 解析扩展类，/lib/ext/*
	loader.parseExtLoader(libPath)
	// 解析用户类
	loader.parseUserLoader(cpOption)
	return loader
}

// 解析启动类
// 记录jre/lib下的所有jar包
// 返回jre/lib路径
func (l *Loader) parseBootLoader(jreOption string) string {
	// 查找jre目录路径
	jreDir := getJreDir(jreOption)

	jreLibPath := filepath.Join(jreDir, "lib", "*")
	// 设置启动类加载器
	l.bootLoader = newWildcardEntry(jreLibPath)

	return filepath.Join(jreDir, "lib")
}

// libPath是在jre/lib的具体路径
// 可以使用在调用 parseBootLoader 后的返回值
func (l *Loader) parseExtLoader(libPath string) {
	// 拼接lib/ext/* 目录 , 然后创建wildcardEntry 对象
	jreExtPath := filepath.Join(libPath, "ext", "*")

	// 设置扩展类加载器
	l.extLoader = newWildcardEntry(jreExtPath)
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
func (l *Loader) parseUserLoader(cption string) {
	// 如果没有用户没有输入 -cp ,默认当前目录为用户目录
	if cption == "" {
		cption = "."
	}
	// 设置应用类加载器
	l.userLoader = newEntry(cption)
}

/*
在classpath 中查找 Class文件
*/
func (l Loader) LoadClass(classpath string) ([]byte, Entry, error) {
	// 拼接类名
	className := classpath + ".class"

	// 在启动加载器中加载类
	if data, entry, err := l.bootLoader.readClass(className); err == nil {
		return data, entry, err
	}

	// 在应用加载器中加载类
	if data, entry, err := l.extLoader.readClass(className); err == nil {
		return data, entry, err
	}

	// 在用户加载器中加载类
	return l.userLoader.readClass(className)

}

/*
toString()
打印用户类得路径
*/
func (l Loader) String() string {
	return l.userLoader.String()
}
