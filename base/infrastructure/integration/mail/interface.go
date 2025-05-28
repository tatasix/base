package mail

import "context"

type MailInterface interface {
	SendNotice(ctx context.Context) error
}
