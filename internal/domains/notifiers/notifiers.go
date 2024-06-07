package notifiers

import (
	"pantori/internal/auth"
	"pantori/internal/domains/goods"
	"pantori/internal/domains/notifiers/core"
	"pantori/internal/domains/notifiers/infra"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() core.ServicePort {
	numWorkers := LoadWorkerNumFromEnv()
	users := auth.NewInternal()
	goods := goods.NewInternal()
	email := infra.NewEmailProvider()

	return core.NewService(goods, users, email, numWorkers)
}

func LoadWorkerNumFromEnv() int {
	viper.BindEnv("notifier_worker_num", "NOTIFIER_WORKER_NUM")
	if viper.IsSet("notifier_worker_num") {
		return viper.GetInt("notifier_worker_num")
	}

	log.Panic().Stack().Err(errors.New("jwt key undefined")).Msg("")
	return 0
}
