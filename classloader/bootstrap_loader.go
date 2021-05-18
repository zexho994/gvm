package classloader

import "path/filepath"

type BootStrapLoader struct {
	path string
	jars []string
}

func newBootStrapLoader(path string) *BootStrapLoader {
	jrePath := getJreDirPath(path)
	libDir := filepath.Join(jrePath, "lib", "*")
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
