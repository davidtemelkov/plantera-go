package api

import (
	"github.com/davidtemelkov/plantera-go/assets"
	"github.com/davidtemelkov/plantera-go/data"
	"github.com/go-chi/chi/v5"
)

func SetUpRoutes() *chi.Mux {
	r := chi.NewRouter()
	assets.Mount(r)

	r.Get("/", handleServeHTML())
	r.Patch("/watered", handlePlantAction(data.WATERED))
	r.Patch("/fertilized", handlePlantAction(data.FERTILIZED))
	r.Patch("/repotted", handlePlantAction(data.REPOTTED))
	r.Post("/plants", handleAddPlant())
	r.Get("/plants/modal/open", handleOpenAddPlantModal())

	return r
}
