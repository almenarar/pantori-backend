package infra

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
)

func createSession() *session.Session {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	if os.Getenv("AWS_ENDPOINT") != "" {
		endpoint := os.Getenv("AWS_ENDPOINT")
		session.Config.Endpoint = &endpoint
	}
	return session
}
