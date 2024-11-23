package myerror

import "fmt"

type MyError struct {
	Msg  string
	Code int
}

func (e MyError) Error() string {
	return fmt.Sprintf("发生了自定义错误，%s，错误码:%d", e.Msg, e.Code)
}

func NewMyError(code int, msg string) *MyError {
	return &MyError{Msg: msg, Code: code}
}
