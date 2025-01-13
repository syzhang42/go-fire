package auth

import (
	"github.com/syzhang42/go-fire/fmtx"
	"github.com/syzhang42/go-fire/log"
)

func StdDebugSlice[T any](perkey string, in []T) {
	log.Debugf("%v:\n%v\n-------------------------\n", perkey, fmtx.FormatSlice(in))
}

func StdDebugMap[T1 comparable, T2 any](perkey string, in map[T1]T2) {
	log.Debugf("%v:\n%v\n-------------------------\n", perkey, fmtx.FormatMap(in))
}

func StdDebugJson[T any](perkey string, in T) {
	log.Debugf("%v:\n%v\n-------------------------\n", perkey, fmtx.FormatOtherJson(in))
}
