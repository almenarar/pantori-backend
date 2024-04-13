package infra_test

import (
	"pantori/internal/auth/core"
	"pantori/internal/auth/infra"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamoDB(t *testing.T) {
	t.Run("DynamoDB Auth Integration Test", func(t *testing.T) {
		assert := assert.New(t)

		dy := infra.NewDynamoDB("pantori-users")

		user := core.User{
			Username:      "int_test_name",
			GivenPassword: "int_test_secret",
			Workspace:     "int_test_workspace",
			Email:         "int_test_email",
			LastSeen:      "int_test_last_seen",
			CreatedAt:     "int_test_created_at",
		}

		err := dy.DeleteUser(user)
		assert.Nil(err)

		output, err := dy.GetUser("int_test_name")
		assert.Equal(core.User{}, output)
		assert.IsType(&infra.ErrUserNotFound{}, err)

		err = dy.CreateUser(user)

		assert.Nil(err)

		output, err = dy.GetUser("int_test_name")

		assert.Nil(err)
		assert.Equal(user.Username, output.Username)
		assert.Equal(user.GivenPassword, output.ActualPassword)
		assert.Equal(user.Workspace, output.Workspace)

		err = dy.DeleteUser(user)
		assert.Nil(err)

	})
}
