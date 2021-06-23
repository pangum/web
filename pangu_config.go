package web

type panguConfig struct {
	// Http配置
	Http struct {
		// 客户端配置
		Server config `json:"server" yaml:"server" validate:"required"`
	} `json:"http" yaml:"http" validate:"required"`
}
