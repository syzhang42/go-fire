package auth

import (
	"fmt"

	"github.com/syzhang42/go-fire/log"
)

// if error != nil, panic("pre:err")
func Must(err error, pre ...string) {
	if err != nil {
		switch len(pre) {
		case 0:
			panic(err)
		case 1:
			panic(fmt.Sprintf("%v:%v", pre[0], err))
		default:
			panic(fmt.Sprintf("%+v:%v", pre, err))
		}
	}
}

// if error != nil, stdlog.Warnf(err)
func LogWarn(err error, pre ...string) {
	if err != nil {
		switch len(pre) {
		case 0:
			log.Warn(err)
		case 1:
			log.Warnf("%v:%v", pre[0], err)
		default:
			log.Warnf("%+v:%v", pre, err)
		}
	}
}
