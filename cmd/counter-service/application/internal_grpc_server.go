package application

import (
	xservers "github.com/syth0le/gopnik/servers"
)

func (a *App) newInternalGRPCServer(env *env) *xservers.GRPCServer {
	server := xservers.NewGRPCServer(
		a.Config.InternalGRPCServer,
		a.Logger,
		xservers.GRPCWithServerName("internal grpc api"),
	)

	return server
}
