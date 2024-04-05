package categories

import (
	"pantori/internal/domains/categories/core"
	"pantori/internal/domains/categories/handlers"
	"pantori/internal/domains/categories/infra"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() *handlers.Network {
	var db core.DatabasePort
	table := loadDynamoDBParamsFromEnv()
	db = infra.NewDynamoDB(table)

	ut := infra.NewUtils()

	service := core.NewService(db, ut)

	return handlers.NewNetwork(service)
}

func loadDynamoDBParamsFromEnv() infra.DynamoParams {
	var params infra.DynamoParams
	if viper.IsSet("aws.categories_table") {
		params.CategoriesTable = viper.GetString("aws.categories_table")
	} else {
		log.Panic().Stack().Err(errors.New("aws.categories_table not set")).Msg("")
	}

	if viper.IsSet("aws.categories_table_index") {
		params.CategoriesTableIndex = viper.GetString("aws.categories_table_index")
	} else {
		log.Panic().Stack().Err(errors.New("aws.categories_table_index not set")).Msg("")
	}
	return params
}
