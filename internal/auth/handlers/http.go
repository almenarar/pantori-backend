package handlers

import (
	"pantori/internal/auth/core"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Network struct {
	svc core.ServicePort
}

func NewNetwork(svc core.ServicePort) *Network {
	return &Network{
		svc: svc,
	}
}

// PingExample godoc
// @Summary Login with username and password
// @Description Endpoint used to login API User
// @Tags Auth
// @Accept json
// @Produce json
// @Param UserLogin body authcore.UserLogin true "UserLogin"
// @Success 200 {string} jwt
// @Router /login [post]
func (net *Network) Login(ctx *gin.Context) {
	userLogin := core.UserLogin{}
	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		log.Error().Stack().Msg(err.Error())
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "username or password field are empty",
			},
		)
		return
	}

	token, err := net.svc.Authenticate(
		core.User{
			Username:      userLogin.Username,
			GivenPassword: userLogin.Password,
		},
	)
	if err != nil {
		statusCode := defineHTTPStatus(err)
		ctx.JSON(
			statusCode,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		token,
	)
}
