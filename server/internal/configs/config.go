package configs

import "simpletools/lib/logger"

type ServerConfig struct {
	Host         string        `json:"host"`          // 本机监听地址 IP:PORT
	RemoteAddr   string        `json:"remote_addr"`   // 远端请求地址 URL
	Platform     string        `json:"platform"`      // 平台
	ShutdownWait int           `json:"shutdown_wait"` // 关闭等待时间
	Debug        bool          `json:"debug"`         // 是否调试模式
	Logger       logger.Config `json:"logger"`
}
