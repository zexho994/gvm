package heap

import (
	"../../classpath"
	"fmt"
)

type ClassLoader struct {
	cp        *classpath.Classpath
	classpath map[string]*Class //loaded classed
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {

}

func (self *ClassLoader) LoadClass(name string) *Class {

}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:        cp,
		classpath: make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class
	}
	return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[LOADED %s from %s]\n", name, entry)
	return class
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException:" + name)
	}
	return data, entry
}
