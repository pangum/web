package web

type config struct {
	// 绑定监听主机
	Host string `json:"host" yaml:"host" xml:"host" toml:"host"`
	// 绑定监听端口
	Port int `default:"9000" json:"port" yaml:"port" xml:"port" toml:"port" validate:"required"`
	// 域名
	Domain string `json:"domain" yaml:"domain" xml:"domain" toml:"domain"`
	// 跨域配置
	Cros crosConfig `json:"cros" yaml:"cros" xml:"cros" toml:"cros" validate:"structonly"`
}
