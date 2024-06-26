package redis

import (
	"context"
	"fmt"

	xerrors "github.com/syth0le/gopnik/errors"
	"go.uber.org/zap"
)

type ClientMock struct {
	Logger *zap.Logger
}

func (c *ClientMock) HSet(ctx context.Context, hasTTL bool, key string, values ...any) error {
	c.Logger.Debug("hset through cache mock")
	return nil
}

func (c *ClientMock) HGet(ctx context.Context, key string, field string) (string, error) {
	c.Logger.Debug("hget through cache mock")
	return "", xerrors.WrapNotFoundError(fmt.Errorf("cannot find smth in cache mock"), "not found in cache")
}

func (c *ClientMock) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	c.Logger.Debug("hget through cache mock")
	return nil, xerrors.WrapNotFoundError(fmt.Errorf("cannot find smth in cache mock"), "not found in cache")
}

func (c *ClientMock) HIncr(ctx context.Context, key, field string) error {
	c.Logger.Debug("hincr through cache mock")
	return nil
}

func (c *ClientMock) Delete(ctx context.Context, keys ...string) error {
	c.Logger.Debug("del through cache mock")
	return nil
}

func (c *ClientMock) Close() error {
	return nil
}
