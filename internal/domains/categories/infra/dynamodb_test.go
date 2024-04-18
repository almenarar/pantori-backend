package infra_test

import (
	"pantori/internal/domains/categories/core"
	"pantori/internal/domains/categories/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamoDB(t *testing.T) {
	t.Run("DynamoDB Category Integration Test", func(t *testing.T) {
		assert := assert.New(t)

		dy := infra.NewDynamoDB(infra.DynamoParams{
			CategoriesTable:      "pantori-categories",
			CategoriesTableIndex: "WorkspaceIndex",
		})

		category := core.Category{
			ID:        "foo",
			Name:      "dinner",
			Color:     "blue",
			Workspace: "test",
		}

		output, err := dy.ListItemsByWorkspace(category.Workspace)
		assert.Nil(err)
		assert.Len(output, 0)

		err = dy.CreateItem(category)
		assert.Nil(err)

		output, err = dy.ListItemsByWorkspace(category.Workspace)
		assert.Nil(err)
		assert.Len(output, 1)
		assert.Equal(output[0].Name, category.Name)
		assert.Equal(output[0].Color, category.Color)
		assert.Equal(output[0].Workspace, category.Workspace)

		err = dy.DeleteItem(output[0])
		assert.Nil(err)
	})
}
