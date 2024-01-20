package auth

import (
	core "pantori/internal/auth/core"
	hdl "pantori/internal/auth/handlers"
	infra "pantori/internal/auth/infra"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() *hdl.Network {
	jwt_key := loadKeyFromEnv()

	crypto := infra.NewCryptography(jwt_key)
	db := infra.NewMemory()

	service := core.NewService(crypto, db)

	return hdl.NewNetwork(service)
}

// it is duplicated from middlewares.go
// i did't found a good way to only do this once
// routes object don't need to know about jwt key env
// main could call a env object to load everything and pass to routes
// but should routes receive a super parameter store and distribute?
// should main create everything and pass them ready to routes?
// i choose to make each module call what they need and keep related things close
func loadKeyFromEnv() string {
	viper.BindEnv("jwt_key", "JWT_KEY")
	if viper.IsSet("jwt_key") {
		return viper.GetString("jwt_key")
	}

	log.Panic().Stack().Err(errors.New("jwt key undefined")).Msg("")

	return ""
}
