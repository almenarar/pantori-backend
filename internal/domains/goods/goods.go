package goods

import (
	core "pantori/internal/domains/goods/core"
	hdl "pantori/internal/domains/goods/handlers"
	infra "pantori/internal/domains/goods/infra"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() *hdl.Network {
	conn := loadConnFromEnv()
	db := infra.NewMySQL(conn)

	params := loadUnsplashParamsFromEnv()
	image := infra.NewUnsplash(params)

	service := core.NewService(db, image)

	return hdl.NewNetwork(service)
}

func loadConnFromEnv() string {
	viper.BindEnv("mysql_conn", "MYSQL_CONN")
	if viper.IsSet("mysql_conn") {
		return viper.GetString("mysql_conn")
	}
	log.Panic().Stack().Err(errors.New("mysql_conn not set")).Msg("")

	return ""
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
