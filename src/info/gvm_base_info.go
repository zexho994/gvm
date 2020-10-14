package info

const version = "v1.9.0"
const author = "zexho"

type gvm struct {
	author  string
	version string
}

func GvmInfo() gvm {
	return gvm{
		author:  author,
		version: version,
	}
}

func (gvm gvm) Version() string {
	return gvm.version
}

func (gvm gvm) Author() string {
	return gvm.author
}
