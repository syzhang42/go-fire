package errorx

import (
	"fmt"

	"github.com/pkg/errors"
)

type errCode int

type _error struct {
	code  int
	desc  string
	cause error
}

var code2Error map[errCode]*_error = map[errCode]*_error{}

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
		code2Error = make(map[errCode]*_error, 0)
	}
	if _, ok := code2Error[errCode(code)]; ok {
		panic("not use same code")
	}
	newError := &_error{
		code: code,
		desc: desc,
	}
	code2Error[errCode(code)] = newError
	return errCode(code)
}
func GetAllRegCodes() map[errCode]*_error {
	return code2Error
}
func (ec errCode) Code() int { return int(ec) }

func (ec errCode) ToError() *_error {
	return code2Error[ec]
}

func (ec errCode) WithError(err error) *_error {
	return code2Error[ec].WithError(err)
}

func (ec errCode) WithMessage(msg string) *_error {
	return code2Error[ec].WithMessage(msg)
}

func (ec errCode) WithMessagef(format string, args ...interface{}) *_error {
	return code2Error[ec].WithMessagef(format, args...)
}

func (e *_error) WithError(err error) *_error {
	if e.cause == nil {
		e.cause = err
	} else {
		e.cause = errors.WithMessage(e.cause, err.Error())
	}
	return e
}

func (e *_error) WithMessage(msg string) *_error {
	if e.cause == nil {
		e.cause = errors.New(msg)
	} else {
		e.cause = errors.WithMessage(e.cause, msg)
	}
	return e
}

func (e *_error) WithMessagef(format string, args ...interface{}) *_error {
	if e.cause == nil {
		e.cause = errors.Errorf(format, args...)
	} else {
		e.cause = errors.WithMessagef(e.cause, format, args...)
	}
	return e
}

func (e *_error) Error() string {
	if e.cause == nil {
		return fmt.Sprintf("[Code:%v, Desc:'%v']", e.code, e.desc)
	}
	return fmt.Sprintf("[Code:%v, Msg:'%v']", e.code, e.desc) + ": " + e.cause.Error()
}

func (e *_error) Cause() error {
	return e.cause
}
