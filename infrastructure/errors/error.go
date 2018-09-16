package errors

import "fmt"

type ErrorCmd struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Errors []ErrorCmd

func (e *ErrorCmd) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

var (
	ErrorCmdRequest  = ErrorCmd{Code: 400, Message: "http request error"}
	ErrorCmdResponse = ErrorCmd{Code: 400, Message: "http response error"}
	ErrorCmdArray    = ErrorCmd{Code: 4001, Message: "json array error"}
)
