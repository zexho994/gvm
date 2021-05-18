package loader

import (
	"github.com/zouzhihao-994/gvm/config"
	"path/filepath"
)

type ExtensionLoader struct {
	path string
	jars []string
}

func newExtensionLoader() *ExtensionLoader {
	extDir := filepath.Join(config.LibJarDir, "lib", "ext")
	config.ExtJarDir = extDir
	el := ExtensionLoader{path: extDir}
	el.jars = jars(filepath.Join(extDir, "*"))
	return &el
}

func (loader *ExtensionLoader) AddZip(s string) {
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
