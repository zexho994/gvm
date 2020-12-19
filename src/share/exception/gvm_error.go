package exception

type GvmError struct {
	Msg string
}

func (e GvmError) Error() string {
	return "[gvm]" + e.Msg
}
