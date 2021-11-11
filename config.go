package web

type config struct {
	// 绑定监听主机
	Host string `json:"host" yaml:"host" xml:"host" toml:"host"`
	// 绑定监听端口
	Port int `default:"9000" json:"port" yaml:"port" xml:"port" toml:"port" validate:"required"`

	// 代理配置
	Proxy *proxy `json:"proxy" yaml:"proxy" xml:"proxy" toml:"proxy" validate:"omitempty,structonly"`
	// 跨域配置
	Cros *cros `json:"cros" yaml:"cros" xml:"cros" toml:"cros" validate:"omitempty,structonly"`
}
