package internalapi

import (
	"context"
	"fmt"
	"net/http"

	xerrors "github.com/syth0le/gopnik/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/syth0le/counter-service/internal/model"
	"github.com/syth0le/counter-service/internal/service/counter"
	inpb "github.com/syth0le/counter-service/proto/internalapi"
)

type CounterHandler struct {
	inpb.UnimplementedCounterServiceServer

	Service counter.Service
}

func (h *CounterHandler) CreateDialogCounters(ctx context.Context, request *inpb.CreateDialogCountersRequest) (*emptypb.Empty, error) {
	// todo: check if counters exist - not create them (not rewrite)
	participants := make([]model.UserID, len(request.Participants))
	for idx, val := range request.Participants {
		participants[idx] = model.UserID(val)
	}

	err := h.Service.CreateDialogCounters(ctx, &counter.CreateDialogCountersParams{
		DialogID:     model.DialogID(request.DialogId),
		Participants: participants,
	})
	if err != nil {
		return nil, GRPCError(fmt.Errorf("create dialog counters: %w", err))
	}

	return &emptypb.Empty{}, nil
}

func (h *CounterHandler) IncreaseDialogCounters(ctx context.Context, request *inpb.IncreaseDialogCountersRequest) (*emptypb.Empty, error) {
	followersIDs := make([]model.UserID, len(request.FollowersId))
	for idx, val := range request.FollowersId {
		followersIDs[idx] = model.UserID(val)
	}

	err := h.Service.IncreaseDialogCounters(ctx, &counter.IncreaseDialogCountersParams{
		DialogID:     model.DialogID(request.DialogId),
		SenderID:     model.UserID(request.SenderId),
		FollowersIDs: followersIDs,
	})
	if err != nil {
		return nil, GRPCError(fmt.Errorf("increase dialog counters: %w", err))
	}

	return &emptypb.Empty{}, nil
}

func (h *CounterHandler) FlushDialogCountersForUser(ctx context.Context, request *inpb.FlushDialogCountersForUserRequest) (*emptypb.Empty, error) {
	err := h.Service.FlushDialogCountersForUser(ctx, &counter.FlushDialogCountersForUserParams{
		DialogID: model.DialogID(request.DialogId),
		UserID:   model.UserID(request.UserId),
	})
	if err != nil {
		return nil, GRPCError(fmt.Errorf("flush dialog counters for user: %w", err))
	}

	return &emptypb.Empty{}, nil
}

func (h *CounterHandler) GetDialogCounterForUser(ctx context.Context, request *inpb.GetDialogCounterForUserRequest) (*inpb.GetDialogCounterForUserResponse, error) {
	counterModel, err := h.Service.GetDialogCounterForUser(ctx, &counter.GetDialogCounterForUserParams{
		DialogID: model.DialogID(request.DialogId),
		UserID:   model.UserID(request.UserId),
	})
	if err != nil {
		return nil, GRPCError(fmt.Errorf("get dialog counter for user: %w", err))
	}

	return &inpb.GetDialogCounterForUserResponse{Counter: int64(counterModel.Value)}, nil
}

func GRPCError(err error) error {
	resError, ok := xerrors.FromError(err)
	if !ok {
		return err
	}

	switch resError.StatusCode {
	case http.StatusForbidden:
		return status.Error(codes.PermissionDenied, resError.Msg)
	case http.StatusNotFound:
		return status.Error(codes.NotFound, resError.Msg)
	case http.StatusBadRequest:
		return status.Error(codes.InvalidArgument, resError.Msg)
	default:
		return status.Errorf(codes.Internal, resError.Msg)
	}
}
