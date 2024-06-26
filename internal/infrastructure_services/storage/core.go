package storage

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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
	GetDialogCounterForUser(ctx context.Context, dialogID model.DialogID, userID model.UserID) (*model.Counter, error)
	GetUserCounters(ctx context.Context, userID model.UserID) ([]*model.Counter, error)
}

type ServiceImpl struct {
	Client redis.Client
	Logger *zap.Logger
}

//           userID
//       /     |     \
//   dialog1  dialog2  dialog3
//     43        3        0
//

func (s *ServiceImpl) CreateDialogCounters(ctx context.Context, dialogID model.DialogID, participants []model.UserID) error {
	// todo: make tx for operation
	dialogkeyHash, err := makeHash(DialogHashType, dialogID.String())
	if err != nil {
		return fmt.Errorf("make hash: %w", err)
	}

	for _, val := range participants {
		userKeyHash, err := makeHash(UserHashType, val.String())
		if err != nil {
			return fmt.Errorf("make hash: %w", err)
		}

		err = s.Client.HSet(
			ctx,
			true,
			userKeyHash,
			dialogkeyHash,
			0,
		)
		if err != nil {
			return fmt.Errorf("cache set: %w", err)
		}

		s.Logger.Sugar().Debugf("key %s saved in cache", userKeyHash)
	}

	return nil
}

func (s *ServiceImpl) IncreaseDialogCounters(ctx context.Context, dialogID model.DialogID, FollowersIDs []model.UserID) error {
	dialogKeyHash, err := makeHash(DialogHashType, dialogID.String())
	if err != nil {
		return fmt.Errorf("make hash: %w", err)
	}

	for _, val := range FollowersIDs {
		userKeyHash, err := makeHash(UserHashType, val.String())
		if err != nil {
			return fmt.Errorf("make hash: %w", err)
		}

		err = s.Client.HIncr(
			ctx,
			userKeyHash,
			dialogKeyHash,
		)
		if err != nil {
			return fmt.Errorf("cache hincr: %w", err)
		}

	}

	s.Logger.Sugar().Debugf("for dialog %s all fields values incremented in cache", dialogKeyHash)

	return nil
}

func (s *ServiceImpl) FlushDialogCountersForUser(ctx context.Context, dialogID model.DialogID, userID model.UserID) error {
	dialogKeyHash, err := makeHash(DialogHashType, dialogID.String())
	if err != nil {
		return fmt.Errorf("make hash: %w", err)
	}

	userKeyHash, err := makeHash(UserHashType, userID.String())
	if err != nil {
		return fmt.Errorf("make hash: %w", err)
	}

	err = s.Client.HSet(
		ctx,
		true,
		userKeyHash,
		dialogKeyHash,
		0,
	)
	if err != nil {
		return fmt.Errorf("cache set: %w", err)
	}

	s.Logger.Sugar().Debugf("key %s %s flushed in cache", userKeyHash, dialogKeyHash)

	return nil
}

func (s *ServiceImpl) GetDialogCounterForUser(ctx context.Context, dialogID model.DialogID, userID model.UserID) (*model.Counter, error) {
	dialogKeyHash, err := makeHash(DialogHashType, dialogID.String())
	if err != nil {
		return nil, fmt.Errorf("make hash: %w", err)
	}

	userKeyHash, err := makeHash(UserHashType, userID.String())
	if err != nil {
		return nil, fmt.Errorf("make hash: %w", err)
	}

	var res *model.CounterValue
	err = s.Client.HGet(
		ctx,
		userKeyHash,
		dialogKeyHash,
		res,
	)
	if err != nil {
		return nil, fmt.Errorf("cache set: %w", err)
	}

	return &model.Counter{DialogID: dialogID, Value: *res}, nil
}

func (s *ServiceImpl) GetUserCounters(ctx context.Context, userID model.UserID) ([]*model.Counter, error) {
	userKeyHash, err := makeHash(UserHashType, userID.String())
	if err != nil {
		return nil, fmt.Errorf("make hash: %w", err)
	}

	dialogCounters, err := s.Client.HGetAll(ctx, userKeyHash)
	if err != nil {
		return nil, fmt.Errorf("cache set: %w", err)
	}

	counters := make([]*model.Counter, len(dialogCounters))
	idx := 0
	for key, val := range dialogCounters {
		counterInt, err := strconv.Atoi(val)
		if err != nil {
			return nil, xerrors.WrapInternalError(fmt.Errorf("cannot get counters from storage: %w", err))
		}

		counters[idx] = &model.Counter{DialogID: parseHash[model.DialogID](key), Value: model.CounterValue(counterInt)}
	}

	return counters, nil
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

func parseHash[T model.DialogID | model.UserID](hash string) T {
	return T(strings.Split(hash, "-")[0])
}
