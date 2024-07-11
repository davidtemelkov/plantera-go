package api

import (
	"fmt"
	"net/http"

	"github.com/davidtemelkov/plantera-go/data"
	"github.com/davidtemelkov/plantera-go/pages"
)

func NewPlantsHandler() PlantsHandler {
	return PlantsHandler{}
}

type PlantsHandler struct {
}

func (nh PlantsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	plants, err := data.GetAllPlants(r.Context())
	if err != nil {
		// TODO: err handling
		fmt.Println(err.Error())
	}

	pages.Plants(plants).Render(r.Context(), w)
}
