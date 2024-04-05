package infra

import (
	"pantori/internal/domains/categories/core"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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
	client := dynamodb.New(createSession())

	av, err := dynamodbattribute.MarshalMap(Category)
	if err != nil {
		return err
	}

	_, err = client.PutItem(
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
	client := dynamodb.New(createSession())
	result, err := client.Query(
		&dynamodb.QueryInput{
			TableName:              &dy.params.CategoriesTable,
			IndexName:              &dy.params.CategoriesTableIndex,
			KeyConditionExpression: aws.String("Workspace = :Workspace"),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":Workspace": {S: aws.String(workspace)},
			},
		},
	)
	if err != nil {
		return []core.Category{}, err
	}

	response := make([]core.Category, 0)
	for _, item := range result.Items {
		Category := core.Category{}
		err := dynamodbattribute.UnmarshalMap(item, &Category)
		if err != nil {
			return []core.Category{}, err
		}
		response = append(response, Category)
	}
	return response, nil
}

func (dy *dynamo) DeleteItem(category core.Category) error {
	client := dynamodb.New(createSession())
	_, err := client.DeleteItem(
		&dynamodb.DeleteItemInput{
			TableName: aws.String(dy.params.CategoriesTable),
			Key: map[string]*dynamodb.AttributeValue{
				"ID": {
					S: aws.String(category.ID),
				},
				"Workspace": {
					S: aws.String(category.Workspace),
				},
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
