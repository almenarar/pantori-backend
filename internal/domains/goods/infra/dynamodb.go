package goodsinfra

import (
	core "pantori/internal/domains/goods/core"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

type DynamoParams struct {
	GoodsTable      string
	GoodsTableIndex string
}

type dynamo struct {
	params DynamoParams
}

func NewDynamoDB(params DynamoParams) *dynamo {
	return &dynamo{
		params: params,
	}
}

func (dy *dynamo) CreateItem(good core.Good) error {
	good.ID = uuid.New().String()
	good.CreatedAt = time.Now().UTC().Format(time.RFC3339)

	err := dy.putItem(good)
	if err != nil {
		return err
	}
	return nil
}

func (dy *dynamo) EditItem(good core.Good) error {
	good.UpdatedAt = time.Now().UTC().Format(time.RFC3339)

	err := dy.putItem(good)
	if err != nil {
		return err
	}
	return nil
}

func (dy *dynamo) putItem(good core.Good) error {
	client := dynamodb.New(createSession())
	_, err := client.PutItem(
		&dynamodb.PutItemInput{
			TableName: aws.String(dy.params.GoodsTable),
			Item: map[string]*dynamodb.AttributeValue{
				"ID": {
					S: aws.String(good.ID),
				},
				"Name": {
					S: aws.String(good.Name),
				},
				"Category": {
					S: aws.String(good.Category),
				},
				"ImageURL": {
					S: aws.String(good.ImageURL),
				},
				"Workspace": {
					S: aws.String(good.Workspace),
				},
				"Expire": {
					S: aws.String(good.Expire),
				},
				"BuyDate": {
					S: aws.String(good.BuyDate),
				},
				"CreatedAt": {
					S: aws.String(good.CreatedAt),
				},
				"UpdatedAt": {
					S: aws.String(good.UpdatedAt),
				},
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (dy *dynamo) GetItemByID(good core.Good) (core.Good, error) {
	client := dynamodb.New(createSession())
	output, err := client.GetItem(
		&dynamodb.GetItemInput{
			TableName: aws.String(dy.params.GoodsTable),
			Key: map[string]*dynamodb.AttributeValue{
				"ID": {
					S: aws.String(dy.params.GoodsTable),
				},
				"Workspace": {
					S: aws.String("main"),
				},
			},
		},
	)
	if err != nil {
		return core.Good{}, err
	}
	if len(output.Item) == 0 {
		return core.Good{}, nil
	} else {
		err = dynamodbattribute.UnmarshalMap(output.Item, &good)
		if err != nil {
			return good, err
		}
	}
	return good, nil
}

func (dy *dynamo) GetAllItems() ([]core.Good, error) {
	client := dynamodb.New(createSession())
	result, err := client.Query(
		&dynamodb.QueryInput{
			TableName:              &dy.params.GoodsTable,
			IndexName:              &dy.params.GoodsTableIndex,
			KeyConditionExpression: aws.String("Workspace = :Workspace"),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":Workspace": {S: aws.String("main")},
			},
		},
	)
	if err != nil {
		return []core.Good{}, err
	}

	response := make([]core.Good, 0)
	for _, item := range result.Items {
		good := core.Good{}
		err := dynamodbattribute.UnmarshalMap(item, &good)
		if err != nil {
			return []core.Good{}, err
		}
		response = append(response, good)
	}
	return response, nil
}

func (dy *dynamo) DeleteItem(good core.Good) error {
	client := dynamodb.New(createSession())
	_, err := client.DeleteItem(
		&dynamodb.DeleteItemInput{
			TableName: aws.String(dy.params.GoodsTable),
			Key: map[string]*dynamodb.AttributeValue{
				"ID": {
					S: aws.String(good.ID),
				},
				"Workspace": {
					S: aws.String("main"),
				},
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
