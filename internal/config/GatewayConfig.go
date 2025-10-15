package config

import (
	"github.com/go-playground/validator/v10"
)

// 网关配置
type GatewayConfig struct {
	Host                string `yaml:"host,default=0.0.0.0" validate:"required,hostname|ip"`                    // 绑定地址
	Port                int    `yaml:"port" validate:"required,gte=1,lte=65535"`                                // 端口
	LogLevel            string `yaml:"log_level,default=info" validate:"omitempty,oneof=debug info warn error"` // 日志级别 debug,info,warn,error
	MaxConns            int    `yaml:"max_conns,default=0" validate:"omitempty,gte=0"`                          // 最大连接数 0表示不限制
	MaxHeaderBytes      int    `yaml:"max_header_bytes,default=1048576" validate:"omitempty,gte=1"`             // 最大请求头大小
	Timeout             int    `yaml:"timeout,default=60" validate:"omitempty,gte=1"`                           // 超时时间 单位秒
	CpuWarningThreshold int    `yaml:"cpu_warning_threshold,default=80" validate:"omitempty,gte=0,lte=100"`     // CPU警告阈值 单位%
	// RouteMaps            []RouteMap           `yaml:"route_maps" validate:"required,dive"`                                     // 路由表
	ClusterServiceConfig ClusterServiceConfig `yaml:"cluster_service_config" validate:"omitempty,dive"` // 集群服务配置
}

var validate = validator.New()

// Validate 验证 GatewayConfig 字段是否合法
func (c *GatewayConfig) Validate() error {
	return validate.Struct(c)
}
