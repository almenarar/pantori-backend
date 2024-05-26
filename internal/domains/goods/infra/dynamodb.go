package infra

import (
	"pantori/internal/domains/goods/core"

	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

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
	client := dynamodb.NewFromConfig(createConfig())

	av, err := attributevalue.MarshalMap(good)
	if err != nil {
		return err
	}

	_, err = client.PutItem(
		context.TODO(),
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
	client := dynamodb.NewFromConfig(createConfig())
	output, err := client.GetItem(
		context.TODO(),
		&dynamodb.GetItemInput{
			TableName: aws.String(dy.params.GoodsTable),
			Key:       dy.GetKey(good),
		},
	)
	if err != nil {
		return core.Good{}, err
	}

	var foundGood core.Good
	if len(output.Item) == 0 {
		return core.Good{}, nil
	} else {
		err = attributevalue.UnmarshalMap(output.Item, &foundGood)
		if err != nil {
			return good, err
		}
	}
	return foundGood, nil
}

func (dy *dynamo) GetAllItems(workspace string) ([]core.Good, error) {
	attr, err := attributevalue.Marshal(workspace)
	if err != nil {
		panic(err)
	}

	client := dynamodb.NewFromConfig(createConfig())
	result, err := client.Query(
		context.TODO(),
		&dynamodb.QueryInput{
			TableName:                 &dy.params.GoodsTable,
			IndexName:                 &dy.params.GoodsTableIndex,
			KeyConditionExpression:    aws.String("Workspace = :Workspace"),
			ExpressionAttributeValues: map[string]types.AttributeValue{":Workspace": attr},
		},
	)

	if err != nil {
		return []core.Good{}, err
	}

	response := make([]core.Good, 0)
	for _, item := range result.Items {
		good := core.Good{}
		err := attributevalue.UnmarshalMap(item, &good)
		if err != nil {
			return []core.Good{}, err
		}
		response = append(response, good)
	}
	return response, nil
}

func (dy *dynamo) DeleteItem(good core.Good) error {
	client := dynamodb.NewFromConfig(createConfig())
	_, err := client.DeleteItem(
		context.TODO(),
		&dynamodb.DeleteItemInput{
			TableName: aws.String(dy.params.GoodsTable),
			Key:       dy.GetKey(good),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (dy *dynamo) GetKey(good core.Good) map[string]types.AttributeValue {
	id, err := attributevalue.Marshal(good.ID)
	if err != nil {
		panic(err)
	}

	workspace, err := attributevalue.Marshal(good.Workspace)
	if err != nil {
		panic(err)
	}

	return map[string]types.AttributeValue{"ID": id, "Workspace": workspace}
}
