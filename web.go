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

	// 处理跨域
	if nil != server.Cros {
		options = append(options, echox.Cros(server.Cros.Credentials, server.Cros.Origins...))
	}
	// 处理代理
	if nil != server.Proxy {
		options = append(options, echox.Proxy(server.Proxy.Scheme, server.Proxy.Domain, server.Proxy.Port))
	}

	echo = &Echo{
		Echo: echox.New(options...),
	}

	return
}
