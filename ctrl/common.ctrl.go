package ctrl

import "fmt"

const (
	SuccessCode   int = 0
	FailedCode    int = 1
	ParamsErrCode int = 2
)

const (
	SuccessCodeMsg   = "request success"
	FailedCodeMsg    = "request failed"
	ParamsErrCodeMsg = "params error"
)

type Err struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e Err) Error() string {
	return fmt.Sprintf("err! code[%d], message[%s]", e.Code, e.Msg)
}

func NewErr(code int, msg string) Err {
	return Err{
		Code: code,
		Msg:  msg,
	}
}

type Response struct {
	Err
	Result interface{} `json:"result"`
}
