package api

import (
	"fmt"
	"net/http"

	"github.com/davidtemelkov/plantera-go/components"
	"github.com/davidtemelkov/plantera-go/data"
	"github.com/davidtemelkov/plantera-go/pages"
	"github.com/davidtemelkov/plantera-go/plants"
)

func handleServeHTML() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		plants, err := data.GetPlants(r.Context(), true)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		pages.Plants(plants).Render(r.Context(), w)
	}
}

func handlePlantAction(action string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		plantID := r.URL.Query().Get("plantId")
		if plantID == "" {
			http.Error(w, "missing plant plantId", http.StatusBadRequest)
			return
		}

		idAttribute := r.URL.Query().Get("id")
		if idAttribute == "" {
			http.Error(w, "missing id attribute", http.StatusBadRequest)
			return
		}

		err := data.UpdatePlant(r.Context(), plantID, action)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		components.ActionZeroDaysAgo(idAttribute, action).Render(r.Context(), w)
	}
}

func handleAddPlant() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// TODO: better error handling, return bad request if malformed input
		newPlant, err := plants.ParsePlantFromRequest(r)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}

		err = data.InsertPlant(r.Context(), newPlant, data.Db)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		// TODO: Instead of this maybe rerender plants
		fmt.Fprintf(w, "plant added successfully")
	}
}

func handleOpenAddPlantModal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		components.AddPlant().Render(r.Context(), w)
	}
}

func handleKillPlant() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		plantID := r.URL.Query().Get("id")
		if plantID == "" {
			http.Error(w, "missing plant id", http.StatusBadRequest)
			return
		}

		err := data.KillPlant(r.Context(), plantID)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		plants, err := data.GetPlants(r.Context(), true)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		components.PlantGrid(plants, true).Render(r.Context(), w)
	}
}

func handleGetGraveyard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		plants, err := data.GetPlants(r.Context(), false)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		pages.Graveyard(plants).Render(r.Context(), w)
	}
}
