package classfile

import "path/filepath"

type BootStrapLoader struct {
	path string
	jars []string
}

func newBootStrapLoader(path string) (*BootStrapLoader, string) {
	jrePath := getJreDirPath(path)
	libDir := filepath.Join(jrePath, "lib", "*")
	jars := jars(libDir)
	return &BootStrapLoader{path: libDir, jars: jars}, jrePath
}

func (loader *BootStrapLoader) Loading() {

}

func (loader *BootStrapLoader) AddZip(zip string) {
	loader.jars = append(loader.jars, zip)
}

func (loader *BootStrapLoader) Path() string {

	return ""
}
