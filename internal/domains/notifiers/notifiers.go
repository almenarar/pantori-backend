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
	emailAuth := LoadEmailAuthFromEnv()

	users := auth.NewInternal()
	goods := goods.NewInternal()
	email := infra.NewEmailProvider(emailAuth)

	return core.NewService(goods, users, email, numWorkers)
}

func LoadWorkerNumFromEnv() int {
	viper.BindEnv("notifier_worker_num", "NOTIFIER_WORKER_NUM")
	if viper.IsSet("notifier_worker_num") {
		return viper.GetInt("notifier_worker_num")
	}

	log.Panic().Stack().Err(errors.New("NOTIFIER_WORKER_NUM undefined")).Msg("")
	return 0
}

func LoadEmailAuthFromEnv() infra.EmailAuth {
	output := infra.EmailAuth{}

	viper.BindEnv("email_email", "EMAIL_PROVIDER_EMAIL")
	if viper.IsSet("email_email") {
		output.Email = viper.GetString("email_email")
	} else {
		log.Panic().Stack().Err(errors.New("EMAIL_PROVIDER_EMAIL undefined")).Msg("")
	}

	viper.BindEnv("email_password", "EMAIL_PROVIDER_PASSWORD")
	if viper.IsSet("email_password") {
		output.Password = viper.GetString("email_password")
	} else {
		log.Panic().Stack().Err(errors.New("EMAIL_PROVIDER_PASSWORD undefined")).Msg("")
	}
	return output
}
