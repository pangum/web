package web

import (
	`github.com/storezhang/echox/v2`
	`github.com/storezhang/pangu`
)

func newEcho(config *pangu.Config) (echo *Echo, err error) {
	panguConfig := new(panguConfig)
	if err = config.Load(panguConfig); nil != err {
		return
	}
	server := panguConfig.Http.Server
	options := echox.NewOptions(echox.Addr(server.Host, server.Port))

	// 处理跨域
	if 0 != len(server.Cros.Origins) {
		options = append(options, echox.Cros(server.Cros.Credentials, server.Cros.Origins...))
	}
	echo = &Echo{
		Echo: echox.New(options...),
	}

	return
}
