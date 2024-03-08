package application

import (
	"context"
	"fmt"
	"syscall"

	xcloser "github.com/syth0le/gopnik/closer"
	"go.uber.org/zap"

	"github.com/syth0le/counter-service/cmd/counter-service/configuration"
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
	_ = envStruct

	internalGrpcServer := a.newInternalGRPCServer(envStruct)
	a.Closer.AddForce(internalGrpcServer.ForcefullyStop)
	a.Closer.Add(internalGrpcServer.GracefullyStop)

	a.Closer.Run(internalGrpcServer.Run)
	a.Closer.Wait()
	return nil
}

type env struct{}

func (a *App) constructEnv(ctx context.Context) (*env, error) {
	//db, err := postgres.NewStorage(a.Logger, a.Config.Storage)
	//if err != nil {
	//	return nil, fmt.Errorf("new storage: %w", err)
	//}
	//a.Closer.Add(db.Close)

	return &env{}, nil
}
