package sms

import "context"

type SmsInterface interface {
	SendNotice(ctx context.Context) error
}
