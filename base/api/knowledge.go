package api

import (
	"base/interfaces/api/handler/course"
	"base/interfaces/api/handler/learningrecord"
	"net/http"

	"github.com/zeromicro/go-zero/rest"

	"base/infrastructure/svc"
	"base/interfaces/api/handler/order"
	"base/interfaces/api/handler/qa"
	"base/interfaces/api/handler/statistics"
	"base/interfaces/api/handler/tool"
	"base/interfaces/api/handler/user"
	"base/interfaces/api/handler/version"
)

// RegisterHandlers 注册HTTP处理器
func RegisterHandlers(server *rest.Server, svc *svc.ServiceContext) {

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/v1/user/login-by-message",
				Handler: user.LoginByMessageHandler(svc),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v1/user/login-by-token",
				Handler: user.LoginByTokenHandler(svc),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v1/user/login-by-apple",
				Handler: user.LoginByAppleHandler(svc),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v1/user/send-message",
				Handler: user.SendMessageHandler(svc),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/user/backend-user",
				Handler: user.BackendUserHandler(svc),
			},
		},
	)

	// 用户
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{svc.Login},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/v1/user/update-user",
					Handler: user.UpdateUserHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/user/delete-account",
					Handler: user.DeleteAccountHandler(svc),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/user/get-user-info",
					Handler: user.GetUserInfoHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/user/add-course",
					Handler: user.AddCourseHandler(svc),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/user/get-my-course",
					Handler: user.GetMyCourseHandler(svc),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/user/my-favorite-count",
					Handler: user.MyFavoriteCountHandler(svc),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/user/get-notification",
					Handler: user.GetNotificationHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/user/read-notification",
					Handler: user.ReadNotificationHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/user/update-mobile",
					Handler: user.UpdateMobileHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/user/send-verify-code",
					Handler: user.SendVerifyCodeHandler(svc),
				},
			}...,
		),
	)
	//qa 知识问答
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{svc.Login},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/v1/qa/search",
					Handler: qa.SearchHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/qa/get-history",
					Handler: qa.GetHistoryHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/qa/get-history-detail",
					Handler: qa.GetHistoryDetailHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/qa/get-retrieval-docs",
					Handler: qa.GetRetrievalDocsHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/qa/change-collection-status",
					Handler: qa.ChangeCollectionStatusHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/qa/get-collection-info",
					Handler: qa.GetCollectionInfoHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/qa/delete-current-session",
					Handler: qa.DeleteCurrentSessionHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/qa/rename-chat-topic",
					Handler: qa.RenameChatTopicHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/qa/stop-chat",
					Handler: qa.StopChatHandler(svc),
				},
			}...,
		),
	)

	// 课程
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{svc.Login},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/v1/courses/get-by-id",
					Handler: course.GetCourseByIDHandler(svc),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/courses/list",
					Handler: course.CourseListHandler(svc),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/courses/detail",
					Handler: course.CourseDetailHandler(svc),
				},
			}...,
		),
	)

	// 课程管理
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/v1/courses/save-course",
				Handler: course.SaveCourseHandler(svc),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v1/courses/save-course-detail",
				Handler: course.SaveCourseDetailHandler(svc),
			},
		},
	)

	// 学习记录
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{svc.Login},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/v1/learning-record/list",
					Handler: learningrecord.LearningRecordListHandler(svc),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/learning-record/details",
					Handler: learningrecord.LearningRecordDetailsHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/learning-record/update",
					Handler: learningrecord.UpdateLearningRecordHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/learning-record/detail/save",
					Handler: learningrecord.SaveLearningRecordDetailHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/learning-record/detail/update",
					Handler: learningrecord.UpdateLearningRecordDetailHandler(svc),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/learning-record/detail",
					Handler: learningrecord.LearningRecordDetailHandler(svc),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/learning-record/recent",
					Handler: learningrecord.LearningRecordRecentHandler(svc),
				},
			}...,
		),
	)

	// 订单
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{svc.Login},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/v1/order/generate",
					Handler: order.GenerateOrderHandler(svc),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/order/get",
					Handler: order.GetOrderHandler(svc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/order/apple-pay-confirm",
					Handler: order.ApplePayConfirmHandler(svc),
				},
			}...,
		),
	)

	// 订单管理
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/v1/order/list",
				Handler: order.ListOrderHandler(svc),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v1/order/refund",
				Handler: order.RefundOrderHandler(svc),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/notify/payment",
				Handler: order.NotifyPaymentHandler(svc),
			},
		},
	)

	// 工具
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/v1/tool/upload-token",
				Handler: tool.UploadTokenHandler(svc),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/tool/test",
				Handler: tool.TestHandler(svc),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v1/tool/add-notification",
				Handler: tool.AddNotificationHandler(svc),
			},
		},
	)

	// 统计
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/v1/statistics/get-daily-metrics",
				Handler: statistics.GetDailyMetricsHandler(svc),
			},
		},
	)

	// 版本
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/v1/version/get-last-version",
				Handler: version.GetLastVersionHandler(svc),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/version/list",
				Handler: version.ListVersionHandler(svc),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v1/version/save",
				Handler: version.SaveVersionHandler(svc),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v1/version/delete",
				Handler: version.DeleteVersionHandler(svc),
			},
		},
	)
}
