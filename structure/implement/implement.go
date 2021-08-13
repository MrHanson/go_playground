package implement

import "fmt"

type RPCError struct {
	Code    int64
	Message string
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

func NewRPCError(code int64, msg string) error {
	return &RPCError{
		Code:    code,
		Message: msg,
	}
}

func AsErr(err error) error {
	return err
}

func ImplementRun() {
	var rpcErr error = NewRPCError(400, "unknown err") // typecheck1
	err := AsErr(rpcErr)                               // typecheck2
	println(err)
}
