package config

// 集群服务配置
type ClusterServiceConfig struct {
	EnableClusterService bool   `yaml:"enable_cluster_service,default=false"`                                  // 是否启用集群服务
	Host                 string `yaml:"host" validate:"required_if=EnableClusterService true,hostname|ip"`     // 绑定地址 切记请使用本地地址不要公开
	Port                 int    `yaml:"port" validate:"required_if=EnableClusterService true,gte=1,lte=65535"` // 端口
}

// Validate 验证 ClusterServiceConfig 字段是否合法
func (c *ClusterServiceConfig) Validate() error {
	return validate.Struct(c)
}
