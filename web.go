package web

import (
	`github.com/pangum/pangu`
	`github.com/storezhang/echox/v2`
)

func newEcho(config *pangu.Config) (echo *Echo, err error) {
	_panguConfig := new(panguConfig)
	if err = config.Load(_panguConfig); nil != err {
		return
	}

	server := _panguConfig.Http.Server
	options := echox.NewOptions(echox.Addr(server.Host, server.Port))

	// 配置域名
	if `` != server.Domain {
		options = append(options, echox.Domain(server.Domain))
	}

	// 处理跨域
	if 0 != len(server.Cros.Origins) {
		options = append(options, echox.Cros(server.Cros.Credentials, server.Cros.Origins...))
	}
	echo = &Echo{
		Echo: echox.New(options...),
	}

	return
}
