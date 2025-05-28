package api

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"

	"base/infrastructure/svc"
	"base/interfaces/api/handler"
)

// RegisterHandlers 注册HTTP处理器
func RegisterHandlers(server *rest.Server, svc *svc.ServiceContext) {

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/v1/notice",
				Handler: handler.NoticeHandler(svc),
			},
		},
	)

}
