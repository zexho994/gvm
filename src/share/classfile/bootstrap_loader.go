package classfile

type BootStrapLoader struct {
	path string
	jars []string
}

func (loader *BootStrapLoader) Loading() {

}

func (loader *BootStrapLoader) AddZip(zip string) {
	loader.jars = append(loader.jars, zip)
}

func (loader *BootStrapLoader) Path() string {

	return ""
}
