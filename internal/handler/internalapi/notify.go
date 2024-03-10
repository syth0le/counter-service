package internalapi

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/syth0le/counter-service/internal/service/notifier"
	"github.com/syth0le/counter-service/proto/internalapi"
)

type NotifierHandler struct {
	internalapi.UnimplementedCounterServiceServer
	NotifierService notifier.Service
}

func (h *NotifierHandler) NotifyMessage(ctx context.Context, request *internalapi.NotifyMessageRequest) (*emptypb.Empty, error) {
	err := h.NotifierService.NotifyMessage(ctx, &notifier.NotifyMessageParams{
		ReaderID: request.ReaderId,
		SenderID: request.SenderId,
		IsRead:   request.IsRead,
	})
	if err != nil {
		return nil, fmt.Errorf("notify message: %w", err)
	}
	return &emptypb.Empty{}, nil
}

func (h *NotifierHandler) mustEmbedUnimplementedCounterServiceServer() {}
