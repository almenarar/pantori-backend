package infra_test

import (
	"pantori/internal/domains/notifiers/core"
	"pantori/internal/domains/notifiers/infra"
	"testing"

	"github.com/spf13/viper"
)

func TestEmail(t *testing.T) {
	t.Run("Email Integration Test", func(t *testing.T) {
		//assert := assert.New(t)
		viper.BindEnv("email_email", "EMAIL_PROVIDER_EMAIL")
		viper.BindEnv("email_password", "EMAIL_PROVIDER_PASSWORD")

		email := infra.NewEmailProvider(infra.EmailAuth{
			Email:    viper.GetString("email_email"),
			Password: viper.GetString("email_password"),
		})

		email.SendEmail(
			core.User{
				Name:      "James Baxter",
				Workspace: "wkp1",
				Email:     "james@gmail.com",
			},
			core.Report{
				ExpiresToday: []core.Good{{
					Name:   "brownie",
					Expire: "30/06/2024",
				}},
				ExpiresSoon: []core.Good{{
					Name:   "pizza",
					Expire: "30/09/2024",
				}},
				Expired: []core.Good{{
					Name:   "queijo",
					Expire: "10/02/2024",
				}},
			})
	})

}
