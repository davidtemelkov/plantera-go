package data

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const (
	AWS_REGION            = "eu-central-1"
	TABLE_NAME            = "plantsV2"
	TIME_FORMAT           = "2006-01-02T15:04:05"
	TIME_FORMAT_JUST_DATE = "2006-01-02T00:00:00"
	PK                    = "PK"
	SK                    = "SK"
	USER_PREFIX           = "USER#"
	PLANT_PREFIX          = "PLANT#"
	ID                    = "Id"
	NAME                  = "Name"
	FERTILIZED            = "Fertilized"
	WATERED               = "Watered"
	REPOTTED              = "Repotted"
	IMAGE_URL             = "ImageURL"
	ALIVE                 = "Alive"
)

var Db *dynamodb.Client

func NewDynamoDbClient(ctx context.Context) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(AWS_REGION))
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
}
