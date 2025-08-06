package sms

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"

	"base/common/util"
	"base/infrastructure/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type Sms struct {
	SvcCtx *svc.ServiceContext
}

func (s *Sms) SendNotice(ctx context.Context) (err error) {
	logger := logx.WithContext(ctx)
	if len(s.SvcCtx.Config.MSG.To) == 0 {
		logger.Errorf("mobile is not null")
		return errors.New("mobile is not null")
	}

	// 构造请求体
	requestBody := map[string]interface{}{
		"mobile":      s.SvcCtx.Config.MSG.To,
		"relation_id": s.SvcCtx.Config.Name + util.NewSnowflake().String(),
		"template_id": s.SvcCtx.Config.MSG.TemplateId,
	}

	// 从配置中获取 SMS 配置
	smsAppId := s.SvcCtx.Config.MSG.AppId
	smsAppSecret := s.SvcCtx.Config.MSG.AppSecret
	smsDomain := s.SvcCtx.Config.MSG.Domain

	// 生成签名信息
	signature, timestamp, nonce := util.GenerateSignature(
		"POST",
		"/v1/sms/send",
		url.Values{},
		requestBody,
		smsAppId,
		smsAppSecret,
	)

	// 将请求体转换为 JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		logger.Errorf("json marshal error: %v", err)
		return err
	}

	// 设置请求头，包含签名信息
	headers := map[string]string{
		"Content-Type": "application/json",
		"x-app-id":     smsAppId,
		"X-Timestamp":  timestamp,
		"X-Nonce":      nonce,
		"X-Signature":  signature,
	}

	// 发送 POST 请求
	response, err := util.Post(smsDomain, jsonData, headers)
	if err != nil {
		logger.Errorf("SendSms error: %v", err)
		return err
	}
	logger.Infof("SendSms url:%s response:%s err:%+v", smsDomain, string(response), err)

	// 解析响应
	var result SendResult
	if err := json.Unmarshal(response, &result); err != nil {
		logger.Errorf("Parse response error: %v", err)
		return err
	}

	logger.Infof("SendSms response: %+v", result)
	if len(result.SmsResponse) > 0 {
		if result.SmsResponse[0].Status == "success" {
			return nil
		}
	}
	return errors.New("message send error")
}

type SendResult struct {
	SmsResponse []struct {
		SerialNo   string `json:"serial_no"`
		RelationId string `json:"relation_id"`
		Status     string `json:"status"`
		Message    string `json:"message"`
	} `json:"sms_response"`
}
