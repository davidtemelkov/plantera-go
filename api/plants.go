package api

import (
	"fmt"
	"net/http"

	"github.com/davidtemelkov/plantera-go/data"
	"github.com/davidtemelkov/plantera-go/pages"
)

func handleServeHTML() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		plants, err := data.GetAllPlants(r.Context())
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		pages.Plants(plants).Render(r.Context(), w)
	}
}

func handlePlantAction(action string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		plantName := r.URL.Query().Get("name")
		if plantName == "" {
			http.Error(w, "missing plant name", http.StatusBadRequest)
			return
		}

		err := data.UpdatePlant(r.Context(), plantName, action)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%s 0 days ago", action)
	}
}
