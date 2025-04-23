package errorx

import (
	"fmt"

	"github.com/pkg/errors"
)

type errCode int

type Error interface {
	WithError(err error) Error
	WithMessage(msg string) Error
	WithMessagef(format string, args ...interface{}) Error
	Error() string
	Cause() error
	Code() int
	Desc() string
}

type localError struct {
	code  int
	desc  string
	cause error
}

func copyError(in Error) (out Error) {
	out = &localError{
		code:  in.Code(),
		desc:  in.Desc(),
		cause: in.Cause(),
	}
	return
}

var code2Error map[errCode]Error = map[errCode]Error{}

/*
注册一个code码,code码保证全局唯一，否则panic。请在初始化逻辑内调用。

var (

		CODE1 = RegErr(0, "code1 text")
		CODE2 = RegErr(0, "code2 text")
		CODE3 = RegErr(2, "code3 text")
		CODE4 = RegErr(3, "code4 text")
		CODE5 = RegErr(4, "code5 text")
		CODE6 = RegErr(5, "code6 text")
		CODE7 = RegErr(6, "code7 text")
	)
*/
func RegErr(code int, desc string) errCode {
	if code2Error == nil {
		code2Error = make(map[errCode]Error, 0)
	}
	if _, ok := code2Error[errCode(code)]; ok {
		panic("not use same code")
	}
	newError := &localError{
		code: code,
		desc: desc,
	}
	code2Error[errCode(code)] = newError
	return errCode(code)
}
func GetAllRegCodes() map[errCode]Error {
	return code2Error
}
func (ec errCode) Code() int { return int(ec) }

func (ec errCode) ToError() Error {
	return code2Error[ec]
}

func (ec errCode) WithError(err error) Error {
	return copyError(code2Error[ec]).WithError(err)
}

func (ec errCode) WithMessage(msg string) Error {
	return copyError(code2Error[ec]).WithMessage(msg)
}

func (ec errCode) WithMessagef(format string, args ...interface{}) Error {
	return copyError(code2Error[ec]).WithMessagef(format, args...)
}

func (e *localError) WithError(err error) Error {
	if e.cause == nil {
		e.cause = err
	} else {
		e.cause = errors.WithMessage(e.cause, err.Error())
	}
	return e
}

func (e *localError) WithMessage(msg string) Error {
	if e.cause == nil {
		e.cause = errors.New(msg)
	} else {
		e.cause = errors.WithMessage(e.cause, msg)
	}
	return e
}

func (e *localError) WithMessagef(format string, args ...interface{}) Error {
	if e.cause == nil {
		e.cause = errors.Errorf(format, args...)
	} else {
		e.cause = errors.WithMessagef(e.cause, format, args...)
	}
	return e
}

func (e *localError) Error() string {
	if e.cause == nil {
		return fmt.Sprintf("[Code:%v, Desc:'%v']", e.code, e.desc)
	}
	return fmt.Sprintf("[Code:%v, Msg:'%v']", e.code, e.desc) + ": " + e.cause.Error()
}

func (e *localError) Cause() error {
	return e.cause
}

func (e *localError) Code() int {
	return e.code
}

func (e *localError) Desc() string {
	return e.desc
}
