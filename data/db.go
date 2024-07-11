package data

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const (
	AWS_REGION = "eu-central-1"
	TABLE_NAME = "plants"
	PK         = "PK"
	NAME       = "Name"
	FERTILIZED = "Fertilized"
	WATERED    = "Watered"
	REPOTTED   = "Repotted"
	IMAGE_URL  = "ImageURL"
)

var Db *dynamodb.Client

func NewDynamoDbClient(ctx context.Context) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(AWS_REGION))
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
}
