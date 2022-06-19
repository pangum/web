package web

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependence(newEcho)
}
