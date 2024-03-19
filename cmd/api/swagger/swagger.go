package swagger

import "os"

type swagger struct{}

func Swagger() *swagger {
	return &swagger{}
}

func (swg *swagger) Config(isDebugging bool) {
	SwaggerInfo.Title = "Pantori API"
	SwaggerInfo.Description = "Serve commands related to Pantori App: auth, workspaces, items and more."
	SwaggerInfo.Version = "1.0"
	SwaggerInfo.BasePath = "/api"
	SwaggerInfo.Schemes = []string{"http"}

	if isDebugging {
		SwaggerInfo.Host = "localhost:8800"
	} else {
		if os.Getenv("ENV") == "production" {
			SwaggerInfo.Host = "pantori-api.ojuqreda8rlp4.us-east-1.cs.amazonlightsail.com"
		}
	}
}
