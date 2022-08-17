package http

import (
	"artforintrovert_test/internal/config"
	"artforintrovert_test/internal/domain/service"
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Rest struct {
	Ctx      context.Context
	Service  *service.TestService
	Port     string
	BasePath string
}

func New(ctx context.Context, serviceApp *service.TestService, cfg *config.Config) *Rest {
	return &Rest{
		Ctx:      ctx,
		Service:  serviceApp,
		Port:     cfg.Listen.Ports.Main,
		BasePath: cfg.Listen.Paths.Base,
	}
}
func (r *Rest) Start(_ context.Context) error {
	router := chi.NewRouter()
	router.Group(r.serviceRoutes)

	mainRouter := chi.NewRouter()
	mainRouter.Mount(r.BasePath, router)
	return http.ListenAndServe(":"+r.Port, mainRouter)
}
