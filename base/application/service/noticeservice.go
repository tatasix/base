package service

import (
	"base/infrastructure/integration/mail"
	"base/infrastructure/integration/sms"
	"base/infrastructure/svc"
	"context"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

// NoticeService 通知应用服务
type NoticeService struct {
	svcCtx         *svc.ServiceContext
	messageService sms.SmsInterface
	mailService    mail.MailInterface
}

// NewNoticeService 创建通知应用服务
func NewNoticeService(svcCtx *svc.ServiceContext, messageService sms.SmsInterface, mailService mail.MailInterface) *NoticeService {
	return &NoticeService{
		svcCtx:         svcCtx,
		messageService: messageService,
		mailService:    mailService,
	}
}

// SendNotice 发送通知
func (s *NoticeService) SendNotice(ctx context.Context) error {
	logger := logx.WithContext(ctx)
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 发送邮件
	go func() {
		defer wg.Done()
		err := s.mailService.SendNotice(ctx)
		if err != nil {
			logger.Errorf("SendNotice error: %v", err)
			return
		}
		logger.Infof("mail SendNotice success")
	}()

	//发送短信
	go func() {
		defer wg.Done()
		err := s.messageService.SendNotice(ctx)
		if err != nil {
			logger.Errorf("SendNotice error: %v", err)
			return
		}
		logger.Infof("message SendNotice success")
	}()
	wg.Wait()
	return nil
}
