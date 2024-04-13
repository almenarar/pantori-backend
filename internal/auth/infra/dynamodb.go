package infra

import (
	"pantori/internal/auth/core"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type dynamo struct {
	UserTable string
}

func NewDynamoDB(table string) *dynamo {
	return &dynamo{
		UserTable: table,
	}
}

func (dy *dynamo) CreateUser(user core.User) error {
	client := dynamodb.New(createSession())

	dbUser := core.UserDB{
		Username:  user.Username,
		Password:  user.GivenPassword,
		Workspace: user.Workspace,
		Email:     user.Email,
		LastSeen:  user.LastSeen,
		CreatedAt: user.CreatedAt,
	}

	av, err := dynamodbattribute.MarshalMap(dbUser)
	if err != nil {
		return &ErrDataTransformFailed{Inner: err}
	}

	_, err = client.PutItem(
		&dynamodb.PutItemInput{
			TableName: aws.String(dy.UserTable),
			Item:      av,
		},
	)
	if err != nil {
		return &ErrDbOpFailed{Inner: err}
	}
	return nil
}

func (dy *dynamo) DeleteUser(user core.User) error {
	client := dynamodb.New(createSession())
	_, err := client.DeleteItem(
		&dynamodb.DeleteItemInput{
			TableName: aws.String(dy.UserTable),
			Key: map[string]*dynamodb.AttributeValue{
				"Username": {
					S: aws.String(user.Username),
				},
			},
		},
	)
	if err != nil {
		return &ErrDbOpFailed{Inner: err}
	}
	return nil
}

func (dy *dynamo) GetUser(username string) (core.User, error) {
	client := dynamodb.New(createSession())
	output, err := client.GetItem(
		&dynamodb.GetItemInput{
			TableName: aws.String(dy.UserTable),
			Key: map[string]*dynamodb.AttributeValue{
				"Username": {
					S: aws.String(username),
				},
			},
		},
	)
	if err != nil {
		return core.User{}, &ErrDbOpFailed{Inner: err}
	}

	var dbUser core.UserDB
	var user core.User
	if len(output.Item) == 0 {
		return core.User{}, &ErrUserNotFound{}
	} else {
		err = dynamodbattribute.UnmarshalMap(output.Item, &dbUser)
		if err != nil {
			return core.User{}, &ErrDataTransformFailed{Inner: err}
		}
		user = core.User{
			Username:       dbUser.Username,
			ActualPassword: dbUser.Password,
			Workspace:      dbUser.Workspace,
		}
	}
	return user, nil
}

func (dy *dynamo) ListUsers() ([]core.User, error) {
	client := dynamodb.New(createSession())
	result, err := client.Scan(
		&dynamodb.ScanInput{
			TableName: aws.String(dy.UserTable),
		},
	)
	if err != nil {
		return []core.User{}, &ErrDbOpFailed{Inner: err}
	}

	response := make([]core.User, 0)
	for _, item := range result.Items {
		user := core.User{}
		err := dynamodbattribute.UnmarshalMap(item, &user)
		if err != nil {
			return response, &ErrDataTransformFailed{Inner: err}
		}
		response = append(response, user)
	}
	return response, nil
}
