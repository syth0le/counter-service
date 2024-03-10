package notifier

import (
	"context"

	"go.uber.org/zap"
)

type Service interface {
	NotifyMessage(ctx context.Context, req *NotifyMessageParams) error
}

type ServiceImpl struct {
	logger *zap.Logger
	//storage
}

func NewService(logger *zap.Logger) *ServiceImpl {
	return &ServiceImpl{logger: logger}
}

type NotifyMessageParams struct {
	ReaderID string
	SenderID string
	IsRead   bool
}

func (s *ServiceImpl) NotifyMessage(ctx context.Context, req *NotifyMessageParams) error {
	if req.IsRead {
		// storage unmove
		return nil
	}
	// storage move
	return nil
}
