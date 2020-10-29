package testlib

// MyError struct
type MyError struct {
	errNo  int
	errMsg string
}

func (me MyError) Error() string {
	return me.errMsg
}
