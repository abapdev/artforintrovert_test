package http

import (
	"artforintrovert_test/internal/config"
	"artforintrovert_test/internal/domain/service"
	"artforintrovert_test/internal/ports"
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Rest struct {
	Ctx      context.Context
	Service  ports.Service
	Port     string
	BasePath string
}

func New(ctx context.Context, serviceApp *service.TestService, cfg *config.Config) *Rest {
	port := cfg.Listen.Ports.Main
	basePath := cfg.Listen.Paths.Base

	return &Rest{
		Ctx:      ctx,
		Service:  serviceApp,
		Port:     port,
		BasePath: basePath,
	}
}
func (r *Rest) Start() error {
	router := chi.NewRouter()
	router.Group(r.serviceRoutes)

	mainRouter := chi.NewRouter()
	mainRouter.Mount(r.BasePath, router)
	return http.ListenAndServe(":"+r.Port, mainRouter)
}
