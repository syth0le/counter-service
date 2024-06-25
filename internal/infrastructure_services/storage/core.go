package storage

import (
	"context"
	"fmt"

	xerrors "github.com/syth0le/gopnik/errors"
	"go.uber.org/zap"

	"github.com/syth0le/counter-service/internal/clients/redis"
	"github.com/syth0le/counter-service/internal/model"
)

const (
	DialogHashType HashType = "dialog"
	UserHashType   HashType = "user"

	defaultRange = 1000 // TODO: pagination
)

type HashType string

func (h HashType) String() string {
	return string(h)
}

type Service interface {
	CreateDialogCounters(ctx context.Context, dialogID model.DialogID, participants []model.UserID) error
	IncreaseDialogCounters(ctx context.Context, dialogID model.DialogID, FollowersIDs []model.UserID) error
	FlushDialogCountersForUser(ctx context.Context, dialogID model.DialogID, userID model.UserID) error
	GetDialogCounterForUser(ctx context.Context, dialogID model.DialogID, userID model.UserID) (*model.CounterValue, error)
	GetUserCounters(ctx context.Context, dialogID model.DialogID, userID model.UserID) (*model.CounterValue, error)
}

type ServiceImpl struct {
	Client redis.Client
	Logger *zap.Logger
}

func (s *ServiceImpl) CreateDialogCounters(ctx context.Context, dialogID model.DialogID, participants []model.UserID) error {
	keyHash, err := makeHash(DialogHashType, dialogID.String())
	if err != nil {
		return fmt.Errorf("make hash: %w", err)
	}

	values := make(map[string]int, len(participants))
	for _, val := range participants {
		valKeyHash, err := makeHash(UserHashType, val.String())
		if err != nil {
			return fmt.Errorf("make hash: %w", err)
		}

		values[valKeyHash] = 0
	}

	err = s.Client.HSet(
		ctx,
		true,
		keyHash,
		values,
	)
	if err != nil {
		return fmt.Errorf("cache set: %w", err)
	}

	s.Logger.Sugar().Infof("key %s saved in cache", keyHash)

	return nil
}

func (s *ServiceImpl) IncreaseDialogCounters(ctx context.Context, dialogID model.DialogID, FollowersIDs []model.UserID) error {
	keyHash, err := makeHash(DialogHashType, dialogID.String())
	if err != nil {
		return fmt.Errorf("make hash: %w", err)
	}

	for _, val := range FollowersIDs {
		valKeyHash, err := makeHash(UserHashType, val.String())
		if err != nil {
			return fmt.Errorf("make hash: %w", err)
		}

		err = s.Client.HIncr(
			ctx,
			keyHash,
			valKeyHash,
		)
		if err != nil {
			return fmt.Errorf("cache hincr: %w", err)
		}
	}

	s.Logger.Sugar().Debugf("for key %s all fields values incremented in cache", keyHash)

	return nil
}

func (s *ServiceImpl) FlushDialogCountersForUser(ctx context.Context, dialogID model.DialogID, userID model.UserID) error {
	keyHash, err := makeHash(DialogHashType, dialogID.String())
	if err != nil {
		return fmt.Errorf("make hash: %w", err)
	}

	valKeyHash, err := makeHash(UserHashType, userID.String())
	if err != nil {
		return fmt.Errorf("make hash: %w", err)
	}

	err = s.Client.HSet(
		ctx,
		true,
		keyHash,
		valKeyHash,
		0,
	)
	if err != nil {
		return fmt.Errorf("cache set: %w", err)
	}

	s.Logger.Sugar().Infof("key %s saved in cache", keyHash)

	return nil
}

func (s *ServiceImpl) GetDialogCounterForUser(ctx context.Context, dialogID model.DialogID, userID model.UserID) (*model.CounterValue, error) {
	keyHash, err := makeHash(DialogHashType, dialogID.String())
	if err != nil {
		return nil, fmt.Errorf("make hash: %w", err)
	}

	valKeyHash, err := makeHash(UserHashType, userID.String())
	if err != nil {
		return nil, fmt.Errorf("make hash: %w", err)
	}

	var res *model.CounterValue
	err = s.Client.HGet(
		ctx,
		keyHash,
		valKeyHash,
		res,
	)
	if err != nil {
		return nil, fmt.Errorf("cache set: %w", err)
	}

	s.Logger.Sugar().Infof("key %s saved in cache", keyHash)

	return res, nil
}

func (s *ServiceImpl) GetUserCounters(ctx context.Context, dialogID model.DialogID, userID model.UserID) (*model.CounterValue, error) {
	// TODO: ???
	return nil, fmt.Errorf("not implemented")
	// keyHash, err := makeHash(DialogHashType, dialogID.String())
	// if err != nil {
	// 	return nil, fmt.Errorf("make hash: %w", err)
	// }
	//
	// valKeyHash, err := makeHash(UserHashType, userID.String())
	// if err != nil {
	// 	return nil, fmt.Errorf("make hash: %w", err)
	// }
	//
	// var res *model.CounterValue
	// err = s.Client.HGet(
	// 	ctx,
	// 	keyHash,
	// 	valKeyHash,
	// 	res,
	// )
	// if err != nil {
	// 	return nil, fmt.Errorf("cache set: %w", err)
	// }
	//
	// s.Logger.Sugar().Infof("key %s saved in cache", keyHash)
	//
	// return res, nil
}

func makeHash(hashType HashType, key string) (string, error) {
	switch hashType {
	case DialogHashType, UserHashType:
	default:
		return "", xerrors.WrapInternalError(fmt.Errorf("unexpected hash type: %s", hashType))
	}

	if key == "" {
		return "", xerrors.WrapInternalError(fmt.Errorf("key cannot be empty"))
	}

	return fmt.Sprintf("%s-%s", hashType, key), nil
}
