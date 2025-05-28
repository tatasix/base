//go:build wireinject
// +build wireinject

package provider

import (
	"base/application/service"
	"base/infrastructure/svc"

	"github.com/google/wire"
)

// providerSet 定义用户服务相关的依赖提供者
var providerSet = wire.NewSet(
	// 应用服务层
	ServiceProviderSet,
)

// InitializeVersionService 初始化版本服务
func InitializeNoticeService(svcCtx *svc.ServiceContext) *service.NoticeService {
	wire.Build(providerSet)
	return nil // 这个返回值会被wire自动生成的代码替换
}

// 注意：执行wire命令后，会生成wire_gen.go文件
// 命令: go run github.com/google/wire/cmd/wire
