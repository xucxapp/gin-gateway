package config

// 路由表
type RouteMap struct {
	RouteType string     `yaml:"route_type,default=http" validate:"required,oneof=http grpc"` // 路由类型 (http转换到http,grpc转换到grpc)
	UpServer  []UpServer `yaml:"up_server" validate:"required,dive"`                          // 上游服务列表
}

// Validate 验证 RouteMap 字段是否合法
func (c *RouteMap) Validate() error {
	return validate.Struct(c)
}
