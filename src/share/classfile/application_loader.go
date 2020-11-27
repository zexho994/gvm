package classfile

type ApplicationLoader struct {
	path string
}

func NewApplicationLoader(path string) *ApplicationLoader {
	return &ApplicationLoader{path: path}
}

func (apploader *ApplicationLoader) AddZip(s string) {
	panic("implement me")
}

func (apploader *ApplicationLoader) Path() string {
	panic("implement me")
}

func (apploader *ApplicationLoader) Loading() {

}
