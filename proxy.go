package web

import (
	`github.com/storezhang/gox`
)

type proxy struct {
	// 协议
	Scheme gox.URIScheme `default:"https" json:"scheme" yaml:"scheme" xml:"scheme" toml:"scheme" validate:"required,oneof=http https"`
	// 域名
	Domain string `json:"domain" yaml:"domain" xml:"domain" toml:"domain" validate:"required,hostname|ip"`
	// 端口
	Port int `default:"443" json:"port" yaml:"port" xml:"port" toml:"port" validate:"required,max=65535"`
}
