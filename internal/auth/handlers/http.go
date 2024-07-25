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
// @Param UserLogin body core.UserLogin true "UserLogin"
// @Success 200 {string} jwt
// @Router /login [post]
func (net *Network) Login(ctx *gin.Context) {
	var err core.DescribedError

	userLogin := core.UserLogin{}
	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		log.Error().Err(err).Msg("/login")
		missingFields := formatValidationError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":  "missing or invalid fields",
			"fields": missingFields,
		})
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
				"error": err.PublicMessage(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		token,
	)
}

// PingExample godoc
// @Summary Create new user
// @Description Endpoint used to create new API user
// @Tags Auth
// @Accept json
// @Produce json
// @Param CreateUser body core.CreateUser true "CreateUser"
// @Success 200 {string} jwt
// @Router /auth/user [post]
// @Security ApiKeyAuth
func (net *Network) CreateUser(ctx *gin.Context) {
	user := core.CreateUser{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Error().Err(err).Msg("/create-user")
		missingFields := formatValidationError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":  "missing or invalid fields",
			"fields": missingFields,
		})
		return
	}

	err := net.svc.CreateUser(
		core.User{
			Username:      user.Username,
			GivenPassword: user.Password,
			Workspace:     user.Workspace,
			Email:         user.Email,
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
		"done!",
	)
}

// PingExample godoc
// @Summary Delete a user
// @Description Endpoint used to delete a API user
// @Tags Auth
// @Accept json
// @Produce json
// @Param DeleteUser body core.DeleteUser true "DeleteUser"
// @Success 200 {string} jwt
// @Router /auth/user [delete]
// @Security ApiKeyAuth
func (net *Network) DeleteUser(ctx *gin.Context) {
	user := core.DeleteUser{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Error().Err(err).Msg("/delete-user")
		missingFields := formatValidationError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":  "missing or invalid fields",
			"fields": missingFields,
		})
		return
	}

	err := net.svc.DeleteUser(
		core.User{
			Username: user.Username,
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
		"done!",
	)
}

// PingExample godoc
// @Summary List users
// @Description Endpoint used to List all users in database
// @Tags Auth
// @Produce json
// @Success 200 {string} arn
// @Router /auth/user [get]
// @Security ApiKeyAuth
func (net *Network) ListUsers(ctx *gin.Context) {
	output, err := net.svc.ListUsers()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		output,
	)
}
