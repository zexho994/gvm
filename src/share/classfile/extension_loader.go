package classfile

import "path/filepath"

type ExtensionLoader struct {
	path string
	jars []string
}

func NewExtensionLoader(path string) *ExtensionLoader {
	extDir := filepath.Join(path, "lib", "ext", "*")
	el := ExtensionLoader{path: extDir}
	el.jars = jars(el.path)
	return &el
}

func (loader *ExtensionLoader) AddZip(s string) {
	panic("implement me")
}

func (loader *ExtensionLoader) Path() string {
	panic("implement me")
}

func (loader *ExtensionLoader) Loading() {

}
