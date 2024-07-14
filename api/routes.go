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
	router.HandleFunc("/watered", handlePlantAction(data.WATERED))
	router.HandleFunc("/fertilized", handlePlantAction(data.FERTILIZED))
	router.HandleFunc("/repotted", handlePlantAction(data.REPOTTED))
	router.HandleFunc("/plants", handleAddPlant())
	router.HandleFunc("/plants/modal/open", handleOpenAddPlantModal())

	return router
}
