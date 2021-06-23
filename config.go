package web

type config struct {
	// 绑定监听主机
	Host string `json:"host" yaml:"host"`
	// 绑定监听端口
	Port int `default:"9000" json:"port" yaml:"port" validate:"required"`
	// 跨域配置
	Cros crosConfig `json:"cros" yaml:"cros" validate:"structonly"`
}
