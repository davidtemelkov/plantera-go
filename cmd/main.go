package main

import (
	"context"
	"net/http"
	"os"

	"github.com/davidtemelkov/plantera-go/api"
	"github.com/davidtemelkov/plantera-go/data"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		os.Exit(1)
	}

	var ctx = context.Background()

	data.Db, err = data.NewDynamoDbClient(ctx)
	if err != nil {
		os.Exit(1)
	}

	router := api.SetUpRoutes()
	if err := http.ListenAndServe(":8080", router); err != nil {
		os.Exit(1)
	}
}
