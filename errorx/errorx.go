package errorx

import (
	"fmt"

	"github.com/pkg/errors"
)

type errCode int

// 不要自己new该struct, 自行new的不会进行统一管理。
// 该结构体封装了 Error 函数，可以视为 error 统一管理。
type Error struct {
	code  int
	desc  string
	cause error
}

var defaultFailedError = &Error{
	code: -10086,
	desc: "errorx usage error !!!!!!,check your AddErrorx,not use same code",
}

var code2Error map[errCode]*Error = map[errCode]*Error{errCode(-10086): defaultFailedError}

/*
如果你可以保证code唯一，可以不判返回的error.
var (

		CODE1, _ = Add(0, "code1 text")
		CODE2, _ = Add(0, "code2 text")
		CODE3, _ = Add(2, "code3 text")
		CODE4, _ = Add(3, "code4 text")
		CODE5, _ = Add(4, "code5 text")
		CODE6, _ = Add(5, "code6 text")
		CODE7, _ = Add(6, "code7 text")
	)
*/
func Add(code int, desc string) (errCode, error) {
	if code2Error == nil {
		code2Error = make(map[errCode]*Error, 0)
	}
	if _, ok := code2Error[errCode(code)]; ok {
		return errCode(-10086), fmt.Errorf("not use same code")
	}
	newError := &Error{
		code: code,
		desc: desc,
	}
	code2Error[errCode(code)] = newError
	return errCode(code), nil
}
func GetAllRegCodes() map[errCode]*Error {
	return code2Error
}
func (ec errCode) Code() int { return int(ec) }

func (ec errCode) ToError() *Error {
	return code2Error[ec]
}

func (ec errCode) WithError(err error) *Error {
	return code2Error[ec].WithError(err)
}

func (ec errCode) WithMessage(msg string) *Error {
	return code2Error[ec].WithMessage(msg)
}

func (ec errCode) WithMessagef(format string, args ...interface{}) *Error {
	return code2Error[ec].WithMessagef(format, args...)
}

func (e *Error) WithError(err error) *Error {
	if e.cause == nil {
		e.cause = err
	} else {
		e.cause = errors.WithMessage(e.cause, err.Error())
	}
	return e
}

func (e *Error) WithMessage(msg string) *Error {
	if e.cause == nil {
		e.cause = errors.New(msg)
	} else {
		e.cause = errors.WithMessage(e.cause, msg)
	}
	return e
}

func (e *Error) WithMessagef(format string, args ...interface{}) *Error {
	if e.cause == nil {
		e.cause = errors.Errorf(format, args...)
	} else {
		e.cause = errors.WithMessagef(e.cause, format, args...)
	}
	return e
}

func (e *Error) Error() string {
	if e.cause == nil {
		return fmt.Sprintf("[Code:%v, Desc:'%v']", e.code, e.desc)
	}
	return fmt.Sprintf("[Code:%v, Msg:'%v']", e.code, e.desc) + ": " + e.cause.Error()
}

func (e *Error) Cause() error {
	return e.cause
}
