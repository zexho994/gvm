package classloader

import (
	"github.com/zouzhihao-994/gvm/config"
	"path/filepath"
)

type BootStrapLoader struct {
	path string
	jars []string
}

func newBootStrapLoader() *BootStrapLoader {
	jrePath := getJreDirPath()
	libDir := filepath.Join(jrePath, "lib", "*")
	config.LibJarDir = jrePath

	jars := jars(libDir)
	return &BootStrapLoader{path: jrePath, jars: jars}
}

func (loader *BootStrapLoader) Loading(fileName string) []byte {
	for _, jar := range loader.jars {
		if data, _ := LoadingFromZip(fileName, jar); data != nil {
			return data
		}
	}
	return nil
}

func (loader *BootStrapLoader) AddZip(zip string) {
	loader.jars = append(loader.jars, zip)
}
