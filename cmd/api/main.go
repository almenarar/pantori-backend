package main

import (
	"pantori/cmd/api/routes"
	"pantori/cmd/api/swagger"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/viper"
)

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	swg := swagger.Swagger()
	swg.Config(gin.IsDebugging())

	viper.SetConfigType("json")
	viper.SetConfigFile("/go/bin/config.json")
	viper.ReadInConfig()

	routes := routes.New()
	routes.Expose()
}
