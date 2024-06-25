package application

import (
	xservers "github.com/syth0le/gopnik/servers"

	"github.com/syth0le/counter-service/internal/handler/internalapi"
	inpb "github.com/syth0le/counter-service/proto/internalapi"
)

func (a *App) newInternalGRPCServer(env *env) *xservers.GRPCServer {
	server := xservers.NewGRPCServer(
		a.Config.InternalGRPCServer,
		a.Logger,
		xservers.GRPCWithServerName("internal grpc api"),
	)

	inpb.RegisterCounterServiceServer(server.Server, &internalapi.CounterHandler{Service: env.counterService})

	return server
}
