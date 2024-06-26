package application

import (
	"context"
	"fmt"
	"syscall"

	xclients "github.com/syth0le/gopnik/clients"
	xcloser "github.com/syth0le/gopnik/closer"
	"go.uber.org/zap"

	"github.com/syth0le/counter-service/cmd/counter/configuration"
	"github.com/syth0le/counter-service/internal/clients/auth"
	"github.com/syth0le/counter-service/internal/clients/redis"
	"github.com/syth0le/counter-service/internal/infrastructure_services/storage"
	"github.com/syth0le/counter-service/internal/service/counter"
)

type App struct {
	Config *configuration.Config
	Logger *zap.Logger
	Closer *xcloser.Closer
}

func New(cfg *configuration.Config, logger *zap.Logger) *App {
	return &App{
		Config: cfg,
		Logger: logger,
		Closer: xcloser.NewCloser(logger, cfg.Application.GracefulShutdownTimeout, cfg.Application.ForceShutdownTimeout, syscall.SIGINT, syscall.SIGTERM),
	}
}

func (a *App) Run() error {
	ctx, cancelFunction := context.WithCancel(context.Background())
	a.Closer.Add(func() error {
		cancelFunction()
		return nil
	})

	envStruct, err := a.constructEnv(ctx)
	if err != nil {
		return fmt.Errorf("construct env: %w", err)
	}

	internalGrpcServer := a.newInternalGRPCServer(envStruct)
	a.Closer.AddForce(internalGrpcServer.ForcefullyStop)
	a.Closer.Add(internalGrpcServer.GracefullyStop)

	httpServer := a.newHTTPServer(envStruct)
	a.Closer.Add(httpServer.GracefulStop()...)

	a.Closer.Run(httpServer.Run()...)
	a.Closer.Run(internalGrpcServer.Run)
	a.Closer.Wait()
	return nil
}

type env struct {
	authClient     auth.Client
	counterService counter.Service
}

func (a *App) constructEnv(ctx context.Context) (*env, error) {
	authClient, err := a.makeAuthClient(ctx, a.Config.AuthClient)
	if err != nil {
		return nil, fmt.Errorf("make auth client: %w", err)
	}

	redisClient := redis.NewRedisClient(a.Logger, a.Config.Storage) // TODO: move to gopnik
	a.Closer.Add(redisClient.Close)

	storageService := &storage.ServiceImpl{Client: redisClient, Logger: a.Logger}

	return &env{
		authClient:     authClient,
		counterService: counter.NewService(a.Logger, storageService),
	}, nil
}

func (a *App) makeAuthClient(ctx context.Context, cfg configuration.AuthClientConfig) (auth.Client, error) {
	if !cfg.Enable {
		return auth.NewClientMock(a.Logger), nil
	}

	connection, err := xclients.NewGRPCClientConn(ctx, cfg.Conn)
	if err != nil {
		return nil, fmt.Errorf("new grpc conn: %w", err)
	}

	a.Closer.Add(connection.Close)

	return auth.NewAuthImpl(a.Logger, connection), nil
}
