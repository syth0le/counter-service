package redis

import (
	"context"
	"encoding"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	xerrors "github.com/syth0le/gopnik/errors"
	"go.uber.org/zap"

	"github.com/syth0le/counter-service/cmd/counter-service/configuration"
)

const (
	defaultClientName = "social-network"
)

type Client interface {
	HSet(ctx context.Context, hasTTL bool, key string, values ...any) error
	HGet(ctx context.Context, key string, field string, scanTo encoding.BinaryUnmarshaler) error
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HIncr(ctx context.Context, key, field string) error
	Delete(ctx context.Context, keys ...string) error
	Close() error
}

type ClientImpl struct {
	Logger             *zap.Logger
	Client             *redis.Client
	ExpirationDuration time.Duration
	MaxListRange       int64
}

func NewRedisClient(logger *zap.Logger, cfg configuration.RedisConfig) Client {
	if !cfg.Enable {
		return &ClientMock{Logger: logger}
	}

	return &ClientImpl{
		Logger: logger,
		Client: redis.NewClient(&redis.Options{
			Addr:       cfg.Address,
			ClientName: defaultClientName,
			Password:   cfg.Password,
			DB:         cfg.Database,
		}),
		ExpirationDuration: cfg.ExpirationDuration,
		MaxListRange:       cfg.MaxListRange,
	}
}

func (c *ClientImpl) Close() error {
	return c.Client.Close()
}

func (c *ClientImpl) HSet(ctx context.Context, hasTTL bool, key string, values ...any) error {
	err := c.Client.HSet(ctx, key, values).Err()
	if err != nil {
		return err
	}

	if hasTTL {
		c.Client.Expire(ctx, key, c.ExpirationDuration)
	}

	return nil
}

func (c *ClientImpl) HGet(ctx context.Context, key string, field string, scanTo encoding.BinaryUnmarshaler) error {
	resp, err := c.Client.HGet(ctx, key, field).Result()
	if err != nil {
		if err != redis.Nil {
			return xerrors.WrapInternalError(fmt.Errorf("hget error"))
		}

		return xerrors.WrapNotFoundError(err, "not found in cache")
	}

	err = scanTo.UnmarshalBinary([]byte(resp))
	if err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}

	return nil
}

func (c *ClientImpl) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	resp, err := c.Client.HGetAll(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			return nil, xerrors.WrapInternalError(fmt.Errorf("hget error"))
		}

		return nil, xerrors.WrapNotFoundError(err, "not found in cache")
	}

	return resp, nil
}

func (c *ClientImpl) HIncr(ctx context.Context, key, field string) error {
	_, err := c.Client.HIncrBy(ctx, key, field, 1).Result()
	if err != nil {
		return fmt.Errorf("hincr by: %w", err)
	}

	return nil
}

func (c *ClientImpl) Delete(ctx context.Context, keys ...string) error {
	_, err := c.Client.Del(ctx, keys...).Result()
	if err != nil {
		return fmt.Errorf("del keys: %w", err)
	}

	return nil
}
