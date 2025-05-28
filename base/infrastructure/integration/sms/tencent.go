package sms

import (
	"base/common/util"
	"base/infrastructure/svc"
	"context"
	"encoding/json"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type Tencent struct {
	SvcCtx *svc.ServiceContext
}

func (s *Tencent) SendNotice(ctx context.Context) (err error) {
	logger := logx.WithContext(ctx)
	if len(s.SvcCtx.Config.MSG.To) == 0 {
		logger.Errorf("mobile is not null")
		return errors.New("mobile is not null")
	}

	mobiles := s.SvcCtx.Config.MSG.To
	// 构造请求体
	requestBody := map[string]interface{}{
		"mobiles":     mobiles,
		"sign_name":   s.SvcCtx.Config.MSG.SignName,
		"template_id": s.SvcCtx.Config.MSG.TemplateId,
	}

	// 将请求体转换为 JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		logger.Errorf("json marshal error: %v", err)
		return err
	}

	// 设置请求头
	headers := map[string]string{
		"Content-Type": "application/json",
		"X-Auth-Key":   s.SvcCtx.Config.MSG.Key,
	}

	// 发送 POST 请求
	response, err := util.Post(s.SvcCtx.Config.MSG.Url, jsonData, headers)
	if err != nil {
		logger.Errorf("SendSms error: %v", err)
		return err
	}
	logger.Infof("SendSms url:%s response:%+v err:%+v", s.SvcCtx.Config.MSG.Url, response, err)
	// 解析响应
	var result map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		logger.Errorf("Parse response error: %v", err)
		return err
	}

	logger.Infof("SendSms response: %+v", result)

	// 检查响应状态
	if code, ok := result["code"].(float64); ok && code != 0 {
		logger.Errorf("SendSms failed with code: %v", code)
		return errors.New("message send error")
	}

	return nil
}
