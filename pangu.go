package web

import (
	`github.com/storezhang/pangu`
)

func init() {
	if err := pangu.New().Provides(newEcho); nil != err {
		panic(err)
	}
}
