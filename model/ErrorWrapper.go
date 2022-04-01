package model

type ErrorWrap struct {
	Code  int
	Error error
}

func (w ErrorWrap) ErrorMsg() string {
	return w.Error.Error()
}
