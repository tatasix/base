package provider

import (
	"base/application/service"
	"base/infrastructure/integration/mail"
	"base/infrastructure/integration/sms"
	"base/infrastructure/svc"

	"github.com/google/wire"
)

// ServiceProviderSet 服务层依赖提供者集合
var ServiceProviderSet = wire.NewSet(
	ProviderNoticeService,
	ProviderMessageService,
	ProviderMailService,
)

func ProviderNoticeService(svcCtx *svc.ServiceContext, messageService sms.SmsInterface, mailService mail.MailInterface) *service.NoticeService {
	return service.NewNoticeService(svcCtx, messageService, mailService)
}

func ProviderMessageService(svcCtx *svc.ServiceContext) sms.SmsInterface {
	return &sms.Tencent{SvcCtx: svcCtx}
}

func ProviderMailService(svcCtx *svc.ServiceContext) mail.MailInterface {
	return &mail.EmailSender{SvcCtx: svcCtx}
}
