package main

import (
	"context"
	"net/http"
	"os"

	"github.com/davidtemelkov/plantera-go/api"
	"github.com/davidtemelkov/plantera-go/data"
)

func main() {
	var ctx = context.Background()

	var err error
	data.Db, err = data.NewDynamoDbClient(ctx)
	if err != nil {
		os.Exit(1)
	}

	router := api.SetUpRoutes()
	if err := http.ListenAndServe(":8080", router); err != nil {
		os.Exit(1)
	}
}
