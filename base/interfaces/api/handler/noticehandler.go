package handler

import (
	"base/infrastructure/provider"
	"base/infrastructure/svc"
	"net/http"
)

// NoticeHandler 处理添加课程的请求
func NoticeHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		service := provider.InitializeNoticeService(svc)
		err := service.SendNotice(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, "success", http.StatusOK)
	}
}
