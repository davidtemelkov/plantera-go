package data

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Plant struct {
	Name       string `json:"name"`
	Fertilized string `json:"fertilized"`
	Repotted   string `json:"repotted"`
	Watered    string `json:"watered"`
	ImageURL   string `json:"imageUrl"`
}

func InsertPlant(ctx context.Context, plant Plant, Db *dynamodb.Client) error {
	item := map[string]types.AttributeValue{
		PK: &types.AttributeValueMemberS{
			Value: NAME,
		},
		FERTILIZED: &types.AttributeValueMemberS{
			Value: plant.Fertilized,
		},
		WATERED: &types.AttributeValueMemberS{
			Value: plant.Watered,
		},
		REPOTTED: &types.AttributeValueMemberS{
			Value: plant.Repotted,
		},
		IMAGE_URL: &types.AttributeValueMemberS{
			Value: plant.ImageURL,
		},
	}

	putInput := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item:      item,
	}

	_, err := Db.PutItem(ctx, putInput)
	if err != nil {
		return err
	}

	return nil
}

func GetAllPlants(ctx context.Context) ([]Plant, error) {
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(TABLE_NAME),
	}

	result, err := Db.Scan(ctx, scanInput)
	if err != nil {
		return nil, err
	}

	plants := make([]Plant, 0)
	for _, item := range result.Items {
		plant := Plant{}
		if err := attributevalue.UnmarshalMap(item, &plant); err != nil {
			return nil, err
		}
		plants = append(plants, plant)
	}

	return plants, nil
}
