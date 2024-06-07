package goods

import (
	core "pantori/internal/domains/goods/core"
	hdl "pantori/internal/domains/goods/handlers"
	infra "pantori/internal/domains/goods/infra"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func NewInternal() *hdl.Internal {
	return hdl.NewInternal(genService())
}

func New() *hdl.Network {
	return hdl.NewNetwork(genService())
}

func genService() core.ServicePort {
	var db core.DatabasePort
	if viper.IsSet("mode.database") &&
		viper.GetString("mode.database") == "dynamo" {
		table := loadDynamoDBParamsFromEnv()
		db = infra.NewDynamoDB(table)
	} else {
		conn := loadConnFromEnv()
		db = infra.NewMySQL(conn)
	}

	params := loadUnsplashParamsFromEnv()
	image := infra.NewUnsplash(params)

	ut := infra.NewUtils()

	return core.NewService(db, image, ut)
}

func loadConnFromEnv() string {
	viper.BindEnv("mysql_conn", "MYSQL_CONN")
	if viper.IsSet("mysql_conn") {
		return viper.GetString("mysql_conn")
	}
	log.Panic().Stack().Err(errors.New("mysql_conn not set")).Msg("")

	return ""
}

func loadDynamoDBParamsFromEnv() infra.DynamoParams {
	var params infra.DynamoParams
	if viper.IsSet("aws.goods_table") {
		params.GoodsTable = viper.GetString("aws.goods_table")
	} else {
		log.Panic().Stack().Err(errors.New("aws.goods_table not set")).Msg("")
	}

	if viper.IsSet("aws.goods_table_index") {
		params.GoodsTableIndex = viper.GetString("aws.goods_table_index")
	} else {
		log.Panic().Stack().Err(errors.New("aws.goods_table_index not set")).Msg("")
	}
	return params
}

func loadUnsplashParamsFromEnv() infra.UnsplashParams {
	var params infra.UnsplashParams

	viper.BindEnv("unsplash_key", "UNSPLASH_KEY")
	if viper.IsSet("unsplash_key") {
		params.AccessKey = viper.GetString("unsplash_key")
	} else {
		log.Panic().Stack().Err(errors.New("unsplash_key not set")).Msg("")
	}

	if viper.IsSet("unsplash.base_url") {
		params.BaseURL = viper.GetString("unsplash.base_url")
	} else {
		log.Panic().Stack().Err(errors.New("unsplash.base_url not set")).Msg("")
	}

	return params
}
