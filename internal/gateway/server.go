package gateway

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-yaml"
	"github.com/xucxapp/gin-gateway/internal/config"
)

type Server struct {
	Config *config.GatewayConfig
	srv    *http.Server
}

// 加载配置文件
func (s *Server) LoadConfig(configFile string) error {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	cfg := &config.GatewayConfig{}
	if len(data) > 0 {
		if err := yaml.Unmarshal(data, cfg); err != nil {
			return err
		}
	}

	if err := cfg.Validate(); err != nil {
		return err
	}

	s.Config = cfg
	return nil
}

// 启动服务
func (s *Server) Start() {
	router := gin.Default()

	s.srv = &http.Server{
		Addr:           s.Config.Host + ":" + strconv.Itoa(s.Config.Port), // 绑定地址和端口
		Handler:        router.Handler(),                                  // 路由处理
		ReadTimeout:    time.Duration(s.Config.Timeout) * time.Second,     // 读取超时时间
		MaxHeaderBytes: s.Config.MaxHeaderBytes,                           // 最大请求头大小
	}
	go func() {
		// service connections
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no params) by default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务器...")

	s.Stop()
}

// 停止服务
func (s *Server) Stop() error {
	// 关闭服务
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Println("服务器关闭失败:", err)
	}
	log.Println("服务器已关闭")
	return nil
}
