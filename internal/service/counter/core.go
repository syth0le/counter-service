package counter

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/syth0le/counter-service/internal/infrastructure_services/storage"
	"github.com/syth0le/counter-service/internal/model"
)

type Service interface {
	CreateDialogCounters(ctx context.Context, params *CreateDialogCountersParams) error
	IncreaseDialogCounters(ctx context.Context, params *IncreaseDialogCountersParams) error
	FlushDialogCountersForUser(ctx context.Context, params *FlushDialogCountersForUserParams) error
	GetDialogCounterForUser(ctx context.Context, params *GetDialogCounterForUserParams) (*model.Counter, error)
	GetUserCounters(ctx context.Context, params *GetUserCountersParams) ([]*model.Counter, error)
}

type ServiceImpl struct {
	logger *zap.Logger

	storageService storage.Service
}

func NewService(logger *zap.Logger, storageService storage.Service) *ServiceImpl {
	return &ServiceImpl{
		logger:         logger,
		storageService: storageService,
	}
}

type CreateDialogCountersParams struct {
	DialogID     model.DialogID
	Participants []model.UserID
}

func (s *ServiceImpl) CreateDialogCounters(ctx context.Context, params *CreateDialogCountersParams) error {
	err := s.storageService.CreateDialogCounters(ctx, params.DialogID, params.Participants)
	if err != nil {
		return fmt.Errorf("create dialog counters: %w", err)
	}

	return nil
}

type IncreaseDialogCountersParams struct {
	DialogID     model.DialogID
	SenderID     model.UserID
	FollowersIDs []model.UserID
}

func (s *ServiceImpl) IncreaseDialogCounters(ctx context.Context, params *IncreaseDialogCountersParams) error {
	err := s.storageService.IncreaseDialogCounters(ctx, params.DialogID, params.FollowersIDs)
	if err != nil {
		return fmt.Errorf("increase dialog counters: %w", err)
	}

	return nil
}

type FlushDialogCountersForUserParams struct {
	DialogID model.DialogID
	UserID   model.UserID
}

func (s *ServiceImpl) FlushDialogCountersForUser(ctx context.Context, params *FlushDialogCountersForUserParams) error {
	err := s.storageService.FlushDialogCountersForUser(ctx, params.DialogID, params.UserID)
	if err != nil {
		return fmt.Errorf("flush dialog counters for user: %w", err)
	}

	return nil
}

type GetDialogCounterForUserParams struct {
	DialogID model.DialogID
	UserID   model.UserID
}

func (s *ServiceImpl) GetDialogCounterForUser(ctx context.Context, params *GetDialogCounterForUserParams) (*model.Counter, error) {
	counter, err := s.storageService.GetDialogCounterForUser(ctx, params.DialogID, params.UserID)
	if err != nil {
		return nil, fmt.Errorf("get dialog counter for user: %w", err)
	}

	return counter, nil
}

type GetUserCountersParams struct {
	UserID model.UserID
}

func (s *ServiceImpl) GetUserCounters(ctx context.Context, params *GetUserCountersParams) ([]*model.Counter, error) {
	counters, err := s.storageService.GetUserCounters(ctx, params.UserID)
	if err != nil {
		return nil, fmt.Errorf("get user counters: %w", err)
	}

	return counters, nil
}
