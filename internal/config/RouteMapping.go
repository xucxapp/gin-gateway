package config

// RouteMapping 路由映射
type RouteMapping struct {
	Method string `yaml:"method,default=GET" validate:"required,oneof=GET POST PUT DELETE"` // 请求方法
	Path   string `yaml:"path" validate:"required"`                                         // 请求路径
	UpPath string `yaml:"up_path" validate:"required"`                                      // 上游路径
}

// Validate 验证 RouteMapping 字段是否合法
func (c *RouteMapping) Validate() error {
	return validate.Struct(c)
}
