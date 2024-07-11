package api

import (
	"github.com/davidtemelkov/plantera-go/assets"
	"github.com/go-chi/chi/v5"
)

func SetUpRoutes() *chi.Mux {
	router := chi.NewRouter()
	assets.Mount(router)

	router.Handle("/", NewPlantsHandler())

	return router
}
