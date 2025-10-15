package config

// UpServer 上游服务
type UpServer struct {
	Name     string         `yaml:"name" validate:"required"`           // 服务名称 GRPC请填写GRPC服务名称
	Mappings []RouteMapping `yaml:"mappings" validate:"omitempty,dive"` // 路由映射
}

// Validate 验证 UpServer 字段是否合法
func (c *UpServer) Validate() error {
	return validate.Struct(c)
}
