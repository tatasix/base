package mail

import (
	"base/infrastructure/svc"
	"context"
	"crypto/tls"
	"fmt"
	"net/smtp"

	"encoding/base64"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ MailInterface = &EmailSender{}

type EmailSender struct {
	SvcCtx *svc.ServiceContext
}

func (s *EmailSender) SendNotice(ctx context.Context) error {
	logger := logx.WithContext(ctx)

	// 邮件内容
	subject := "系统告警"
	body := fmt.Sprintf(`
		<html>
		<body>
			<h3>您好！</h3>
			<p>您的系统出现告警，请及时处理。</p>
		</body>
		</html>
	`)
	tos := s.SvcCtx.Config.Email.To
	for _, to := range tos {
		err := s.doSend(ctx, subject, body, to)
		if err != nil {
			logger.Errorf("发送邮件失败: %v", err)
			return err
		}
	}

	return nil
}

func (s *EmailSender) doSend(ctx context.Context, to, subject, body string) error {
	logger := logx.WithContext(ctx)

	// 构建邮件头
	headers := make(map[string]string)
	headers["From"] = s.SvcCtx.Config.Email.From
	headers["To"] = to
	headers["Subject"] = "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(subject)) + "?="
	headers["Content-Type"] = "text/html; charset=UTF-8"
	headers["MIME-Version"] = "1.0"

	// 组装邮件内容
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// 发送邮件
	auth := smtp.PlainAuth("", s.SvcCtx.Config.Email.Username, s.SvcCtx.Config.Email.Password, s.SvcCtx.Config.Email.Host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         s.SvcCtx.Config.Email.Host,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", s.SvcCtx.Config.Email.Host, s.SvcCtx.Config.Email.Port), tlsConfig)
	if err != nil {
		logger.Errorf("连接邮件服务器失败: %v", err)
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, s.SvcCtx.Config.Email.Host)
	if err != nil {
		logger.Errorf("创建SMTP客户端失败: %v", err)
		return err
	}
	defer client.Close()

	if err = client.Auth(auth); err != nil {
		logger.Errorf("SMTP认证失败: %v", err)
		return err
	}

	if err = client.Mail(s.SvcCtx.Config.Email.From); err != nil {
		logger.Errorf("设置发件人失败: %v", err)
		return err
	}

	if err = client.Rcpt(to); err != nil {
		logger.Errorf("设置收件人失败: %v", err)
		return err
	}

	w, err := client.Data()
	if err != nil {
		logger.Errorf("准备发送数据失败: %v", err)
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		logger.Errorf("写入邮件内容失败: %v", err)
		return err
	}

	err = w.Close()
	if err != nil {
		logger.Errorf("关闭数据写入失败: %v", err)
		return err
	}
	logger.Info("邮件发送成功")
	return nil

}
