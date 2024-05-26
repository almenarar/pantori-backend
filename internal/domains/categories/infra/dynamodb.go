package infra

import (
	"pantori/internal/domains/categories/core"

	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoParams struct {
	CategoriesTable      string
	CategoriesTableIndex string
}

type dynamo struct {
	params DynamoParams
}

func NewDynamoDB(params DynamoParams) *dynamo {
	return &dynamo{
		params: params,
	}
}

func (dy *dynamo) CreateItem(Category core.Category) error {
	err := dy.putItem(Category)
	if err != nil {
		return err
	}
	return nil
}

func (dy *dynamo) EditItem(Category core.Category) error {
	err := dy.putItem(Category)
	if err != nil {
		return err
	}
	return nil
}

func (dy *dynamo) putItem(Category core.Category) error {
	client := dynamodb.NewFromConfig(createConfig())

	av, err := attributevalue.MarshalMap(Category)
	if err != nil {
		return err
	}

	_, err = client.PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{
			TableName: aws.String(dy.params.CategoriesTable),
			Item:      av,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (dy *dynamo) ListItemsByWorkspace(workspace string) ([]core.Category, error) {
	attr, err := attributevalue.Marshal(workspace)
	if err != nil {
		panic(err)
	}

	client := dynamodb.NewFromConfig(createConfig())
	result, err := client.Query(
		context.TODO(),
		&dynamodb.QueryInput{
			TableName:                 &dy.params.CategoriesTable,
			IndexName:                 &dy.params.CategoriesTableIndex,
			KeyConditionExpression:    aws.String("Workspace = :Workspace"),
			ExpressionAttributeValues: map[string]types.AttributeValue{":Workspace": attr},
		},
	)
	if err != nil {
		return []core.Category{}, err
	}

	response := make([]core.Category, 0)
	for _, item := range result.Items {
		Category := core.Category{}
		err := attributevalue.UnmarshalMap(item, &Category)
		if err != nil {
			return []core.Category{}, err
		}
		response = append(response, Category)
	}
	return response, nil
}

func (dy *dynamo) DeleteItem(category core.Category) error {
	client := dynamodb.NewFromConfig(createConfig())
	_, err := client.DeleteItem(
		context.TODO(),
		&dynamodb.DeleteItemInput{
			TableName: aws.String(dy.params.CategoriesTable),
			Key:       dy.GetKey(category),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (dy *dynamo) GetKey(category core.Category) map[string]types.AttributeValue {
	id, err := attributevalue.Marshal(category.ID)
	if err != nil {
		panic(err)
	}

	workspace, err := attributevalue.Marshal(category.Workspace)
	if err != nil {
		panic(err)
	}

	return map[string]types.AttributeValue{"ID": id, "Workspace": workspace}
}
