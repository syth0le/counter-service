package application

import (
	"github.com/go-chi/chi/v5"
	xservers "github.com/syth0le/gopnik/servers"

	"github.com/syth0le/counter-service/internal/handler/publicapi"
)

func (a *App) newHTTPServer(env *env) *xservers.HTTPServerWrapper {
	return xservers.NewHTTPServerWrapper(
		a.Logger,
		xservers.WithAdminServer(a.Config.AdminServer),
		xservers.WithPublicServer(a.Config.PublicServer, a.publicMux(env)),
	)
}

func (a *App) publicMux(env *env) *chi.Mux {
	mux := chi.NewMux()

	handler := &publicapi.Handler{
		Logger:         a.Logger,
		CounterService: env.counterService,
	}

	mux.Route("/counter", func(r chi.Router) {
		r.Use(env.authClient.AuthenticationInterceptor)

		r.Get("/", handler.GetUserCountersList)
		r.Get("/{dialogID}", handler.GetUserCounter)
	})

	return mux
}
