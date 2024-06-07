package auth

import (
	"pantori/internal/auth/core"
	"pantori/internal/auth/handlers"
	"pantori/internal/auth/infra"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func NewInternal() *handlers.Internal {
	return handlers.NewInternal(genService())
}

func New() *handlers.Network {
	return handlers.NewNetwork(genService())
}

func genService() core.ServicePort {
	jwt_key := LoadKeyFromEnv()
	table := LoadDynamoDBParamsFromEnv()

	crypto := infra.NewCryptography(jwt_key)
	utils := infra.NewUtils()
	db := infra.NewDynamoDB(table)

	return core.NewService(crypto, db, utils)
}

func LoadDynamoDBParamsFromEnv() string {
	var table string
	if viper.IsSet("aws.user_table") {
		table = viper.GetString("aws.user_table")
	} else {
		log.Panic().Stack().Err(errors.New("aws.user_table not set")).Msg("")
	}

	return table
}

// it is duplicated from middlewares.go
// i did't found a good way to only do this once
// routes object don't need to know about jwt key env
// main could call a env object to load everything and pass to routes
// but should routes receive a super parameter store and distribute?
// should main create everything and pass them ready to routes?
// i choose to make each module call what they need and keep related things close
func LoadKeyFromEnv() string {
	viper.BindEnv("jwt_key", "JWT_KEY")
	if viper.IsSet("jwt_key") {
		return viper.GetString("jwt_key")
	}

	log.Panic().Stack().Err(errors.New("jwt key undefined")).Msg("")

	return ""
}
