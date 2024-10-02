package data

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type Plant struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Fertilized string `json:"fertilized"`
	Repotted   string `json:"repotted"`
	Watered    string `json:"watered"`
	ImageURL   string `json:"imageUrl"`
	Alive      bool   `json:"alive"`
}

func InsertPlant(ctx context.Context, plant Plant, Db *dynamodb.Client) error {
	plant.ID = uuid.New().String()

	item := map[string]types.AttributeValue{
		PK: &types.AttributeValueMemberS{
			Value: USER_PREFIX + "test@mail.com",
		},
		SK: &types.AttributeValueMemberS{
			Value: PLANT_PREFIX + plant.ID,
		},
		ID: &types.AttributeValueMemberS{
			Value: plant.ID,
		},
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

func GetPlants(ctx context.Context, isAlive bool) ([]Plant, error) {
	keyConditionExpression := "#pk = :pk AND begins_with(#sk, :sk)"
	expressionAttributeNames := map[string]string{
		"#pk": PK,
		"#sk": SK,
	}
	expressionAttributeValues := map[string]types.AttributeValue{
		":pk": &types.AttributeValueMemberS{
			Value: USER_PREFIX + "test@mail.com",
		},
		":sk": &types.AttributeValueMemberS{
			Value: PLANT_PREFIX,
		},
		":alive": &types.AttributeValueMemberBOOL{
			Value: isAlive,
		},
	}

	filterExpression := "Alive = :alive"

	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(TABLE_NAME),
		KeyConditionExpression:    aws.String(keyConditionExpression),
		ExpressionAttributeNames:  expressionAttributeNames,
		ExpressionAttributeValues: expressionAttributeValues,
		FilterExpression:          aws.String(filterExpression),
	}

	result, err := Db.Query(ctx, queryInput)
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

func UpdatePlant(ctx context.Context, plantID, action string) error {
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
			PK: &types.AttributeValueMemberS{Value: USER_PREFIX + "test@mail.com"},
			SK: &types.AttributeValueMemberS{Value: PLANT_PREFIX + plantID},
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

func KillPlant(ctx context.Context, plantID string) error {
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
			PK: &types.AttributeValueMemberS{Value: USER_PREFIX + "test@mail.com"},
			SK: &types.AttributeValueMemberS{Value: PLANT_PREFIX + plantID},
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
