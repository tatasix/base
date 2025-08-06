package config

import (
	"github.com/zeromicro/go-zero/rest"
)

// Config 应用配置
type Config struct {
	rest.RestConf // REST服务配置

	Redis struct {
		Host string // Redis主机
		Pass string // Redis密码
		Type string // Redis类型
		Tls  bool   // Redis是否启用TLS
	}
	MSG struct {
		To         []string // 发送者
		TemplateId string   // 短信模板
		Url        string   // 短信URL
		Key        string   // API密钥
		SignName   string   // 账号
		Domain     string   // 域名
		AppId      string   // 应用ID
		AppSecret  string   // 应用密钥
	}

	Email struct {
		Host        string
		Port        int
		Username    string
		Password    string
		From        string
		FrontendURL string
		To          []string
	}
}
