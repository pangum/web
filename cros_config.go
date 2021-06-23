package web

type crosConfig struct {
	// 允许访问资源的地址
	Origins []string `json:"origins" yaml:"origins"`
	// 是否可以将对请求的响应暴露给页面
	Credentials bool `default:"true" json:"credentials" yaml:"credentials"`
}
