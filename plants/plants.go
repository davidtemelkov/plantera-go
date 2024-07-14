package plants

import (
	"errors"
	"net/http"
	"time"

	"github.com/davidtemelkov/plantera-go/data"
	"github.com/davidtemelkov/plantera-go/utils"
)

// TODO: Add form validation
func ParsePlantFromRequest(r *http.Request) (data.Plant, error) {
	name := r.FormValue("name")
	watered := r.FormValue("watered")
	repotted := r.FormValue("repotted")
	fertilized := r.FormValue("fertilized")

	wateredFormatted, err := formatDate(watered)
	if err != nil {
		return data.Plant{}, errors.New("watered date error")
	}

	fertilizedFormatted, err := formatDate(fertilized)
	if err != nil {
		return data.Plant{}, errors.New("fertilized date error")
	}

	repottedFormatted, err := formatDate(repotted)
	if err != nil {
		return data.Plant{}, errors.New("repotted date error")
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		return data.Plant{}, errors.New("retrieving file error ")
	}
	defer file.Close()

	imageURL, err := utils.UploadFile(r.Context(), file)
	if err != nil {
		return data.Plant{}, errors.New("upload file error")
	}

	newPlant := data.Plant{
		Name:       name,
		Watered:    wateredFormatted,
		Repotted:   repottedFormatted,
		Fertilized: fertilizedFormatted,
		ImageURL:   imageURL,
	}

	return newPlant, nil
}

func formatDate(dateStr string) (string, error) {
	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", err
	}
	return parsedTime.Format(data.TIME_FORMAT_JUST_DATE), nil
}
