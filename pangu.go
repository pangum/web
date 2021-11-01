package web

import (
	`github.com/pangum/pangu`
)

func init() {
	if err := pangu.New().Provides(newEcho); nil != err {
		panic(err)
	}
}
