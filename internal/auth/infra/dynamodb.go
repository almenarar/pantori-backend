package infra

import (
	"pantori/internal/auth/core"

	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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
	client := dynamodb.NewFromConfig(createConfig())

	dbUser := core.UserDB{
		Username:  user.Username,
		Password:  user.GivenPassword,
		Workspace: user.Workspace,
		Email:     user.Email,
		LastSeen:  user.LastSeen,
		CreatedAt: user.CreatedAt,
	}

	av, err := attributevalue.MarshalMap(dbUser)
	if err != nil {
		return &ErrDataTransformFailed{Inner: err}
	}

	_, err = client.PutItem(
		context.TODO(),
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
	client := dynamodb.NewFromConfig(createConfig())
	_, err := client.DeleteItem(
		context.TODO(),
		&dynamodb.DeleteItemInput{
			TableName: aws.String(dy.UserTable),
			Key:       dy.GetKey(user),
		},
	)
	if err != nil {
		return &ErrDbOpFailed{Inner: err}
	}
	return nil
}

func (dy *dynamo) GetUser(user core.User) (core.User, error) {
	client := dynamodb.NewFromConfig(createConfig())
	output, err := client.GetItem(
		context.TODO(),
		&dynamodb.GetItemInput{
			TableName: aws.String(dy.UserTable),
			Key:       dy.GetKey(user),
		},
	)
	if err != nil {
		return core.User{}, &ErrDbOpFailed{Inner: err}
	}

	var dbUser core.UserDB
	if len(output.Item) == 0 {
		return core.User{}, &ErrUserNotFound{}
	} else {
		err = attributevalue.UnmarshalMap(output.Item, &dbUser)
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
	client := dynamodb.NewFromConfig(createConfig())
	result, err := client.Scan(
		context.TODO(),
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
		err := attributevalue.UnmarshalMap(item, &user)
		if err != nil {
			return response, &ErrDataTransformFailed{Inner: err}
		}
		response = append(response, user)
	}
	return response, nil
}

func (dy *dynamo) GetKey(user core.User) map[string]types.AttributeValue {
	username, err := attributevalue.Marshal(user.Username)
	if err != nil {
		panic(err)
	}

	return map[string]types.AttributeValue{"Username": username}
}
