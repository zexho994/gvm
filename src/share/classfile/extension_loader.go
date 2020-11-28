package classfile

import "path/filepath"

type ExtensionLoader struct {
	path string
	jars []string
}

func NewExtensionLoader(path string) *ExtensionLoader {
	extDir := filepath.Join(path, "lib", "ext")
	el := ExtensionLoader{path: extDir}
	el.jars = jars(filepath.Join(extDir, "*"))
	return &el
}

func (loader *ExtensionLoader) AddZip(s string) {
	panic("implement me")
}

func (loader *ExtensionLoader) Path() string {
	panic("implement me")
}

func (loader *ExtensionLoader) Loading(fileName string) []byte {
	for _, jar := range loader.jars {
		if data, _ := LoadingFromZip(fileName, jar); data != nil {
			return data
		}
	}
	return nil
}
