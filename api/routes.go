package api

import (
	"github.com/davidtemelkov/plantera-go/assets"
	"github.com/davidtemelkov/plantera-go/data"
	"github.com/go-chi/chi/v5"
)

func SetUpRoutes() *chi.Mux {
	router := chi.NewRouter()
	assets.Mount(router)

	router.HandleFunc("/", handleServeHTML())
	router.HandleFunc("/water", handlePlantAction(data.WATERED))
	router.HandleFunc("/fertilize", handlePlantAction(data.FERTILIZED))
	router.HandleFunc("/repot", handlePlantAction(data.REPOTTED))

	return router
}
