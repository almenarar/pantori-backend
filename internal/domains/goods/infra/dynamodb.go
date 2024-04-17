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

	av, err := dynamodbattribute.MarshalMap(good)
	if err != nil {
		return err
	}

	_, err = client.PutItem(
		&dynamodb.PutItemInput{
			TableName: aws.String(dy.params.GoodsTable),
			Item:      av,
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
					S: aws.String(good.ID),
				},
				"Workspace": {
					S: aws.String(good.Workspace),
				},
			},
		},
	)
	if err != nil {
		return core.Good{}, err
	}

	var foundGood core.Good
	if len(output.Item) == 0 {
		return core.Good{}, nil
	} else {
		err = dynamodbattribute.UnmarshalMap(output.Item, &foundGood)
		if err != nil {
			return good, err
		}
	}
	return foundGood, nil
}

func (dy *dynamo) GetAllItems(workspace string) ([]core.Good, error) {
	client := dynamodb.New(createSession())
	result, err := client.Query(
		&dynamodb.QueryInput{
			TableName:              &dy.params.GoodsTable,
			IndexName:              &dy.params.GoodsTableIndex,
			KeyConditionExpression: aws.String("Workspace = :Workspace"),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":Workspace": {S: aws.String(workspace)},
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
					S: aws.String(good.Workspace),
				},
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
