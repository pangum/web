package web

import (
	`github.com/storezhang/gox`
)

type config struct {
	// 协议
	Scheme gox.URIScheme `default:"https" json:"scheme" yaml:"scheme" xml:"scheme" toml:"scheme"`
	// 绑定监听主机
	Host string `json:"host" yaml:"host" xml:"host" toml:"host"`
	// 绑定监听端口
	Port int `default:"9000" json:"port" yaml:"port" xml:"port" toml:"port" validate:"required"`
	// 跨域配置
	Cros cros `json:"cros" yaml:"cros" xml:"cros" toml:"cros" validate:"structonly"`
}
