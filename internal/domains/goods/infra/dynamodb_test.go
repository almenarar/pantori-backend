package infra_test

import (
	"pantori/internal/domains/goods/core"
	infra "pantori/internal/domains/goods/infra"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamoDB(t *testing.T) {
	t.Run("DynamoDB Good Integration Test", func(t *testing.T) {
		assert := assert.New(t)

		dy := infra.NewDynamoDB(infra.DynamoParams{
			GoodsTable:      "pantori-goods",
			GoodsTableIndex: "WorkspaceIndex",
		})

		good := core.Good{
			Name:       "grapes",
			Categories: []string{"fruit"},
			Workspace:  "int_test_workspace",
			CreatedAt:  "int_test_created_at",
		}

		output, _ := dy.GetItemByID(good)
		assert.Equal(core.Good{}, output)

		err := dy.CreateItem(good)
		assert.Nil(err)

		goods, err := dy.GetAllItems(good.Workspace)
		assert.Nil(err)
		assert.Len(goods, 1)

		output, err = dy.GetItemByID(goods[0])
		assert.Nil(err)
		assert.Equal(good.Name, output.Name)
		assert.Equal(good.Workspace, output.Workspace)
		assert.Equal(good.Categories, output.Categories)

		err = dy.DeleteItem(goods[0])
		assert.Nil(err)

	})

}
