package data

import (
	"context"
	"errors"
	"time"

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
	Alive      bool   `json:"alive"`
}

func InsertPlant(ctx context.Context, plant Plant, Db *dynamodb.Client) error {
	item := map[string]types.AttributeValue{
		NAME: &types.AttributeValueMemberS{
			Value: plant.Name,
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
		ALIVE: &types.AttributeValueMemberBOOL{
			Value: true,
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

func GetAllLivingPlants(ctx context.Context) ([]Plant, error) {
	scanInput := &dynamodb.ScanInput{
		TableName:        aws.String(TABLE_NAME),
		FilterExpression: aws.String("Alive = :alive"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":alive": &types.AttributeValueMemberBOOL{Value: true},
		},
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

func UpdatePlant(ctx context.Context, plantName, action string) error {
	if action != WATERED &&
		action != FERTILIZED &&
		action != REPOTTED {
		return errors.New("no such action")
	}

	currentTime := time.Now().UTC().Format(TIME_FORMAT)

	updateExpression := "SET #A = :a"
	expressionAttributeNames := map[string]string{
		"#A": action,
	}
	expressionAttributeValues := map[string]types.AttributeValue{
		":a": &types.AttributeValueMemberS{Value: currentTime},
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]types.AttributeValue{
			NAME: &types.AttributeValueMemberS{Value: plantName},
		},
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeNames:  expressionAttributeNames,
		ExpressionAttributeValues: expressionAttributeValues,
	}

	_, err := Db.UpdateItem(ctx, updateInput)
	if err != nil {
		return err
	}

	return nil
}

func KillPlant(ctx context.Context, plantName string) error {
	updateExpression := "SET #A = :a"
	expressionAttributeNames := map[string]string{
		"#A": ALIVE,
	}
	expressionAttributeValues := map[string]types.AttributeValue{
		":a": &types.AttributeValueMemberBOOL{Value: false},
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]types.AttributeValue{
			NAME: &types.AttributeValueMemberS{Value: plantName},
		},
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeNames:  expressionAttributeNames,
		ExpressionAttributeValues: expressionAttributeValues,
	}

	_, err := Db.UpdateItem(ctx, updateInput)
	if err != nil {
		return err
	}

	return nil
}

func GetAllDeadPlants(ctx context.Context) ([]Plant, error) {
	scanInput := &dynamodb.ScanInput{
		TableName:        aws.String(TABLE_NAME),
		FilterExpression: aws.String("Alive = :alive"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":alive": &types.AttributeValueMemberBOOL{Value: false},
		},
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
