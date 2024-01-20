package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type middlewares struct {
	jwtKey string
}

func New() *middlewares {
	var key string
	viper.BindEnv("jwt_key", "JWT_KEY")
	if viper.IsSet("jwt_key") {
		key = viper.GetString("jwt_key")
	} else {
		log.Panic().Stack().Err(errors.New("jwt key undefined")).Msg("")
	}
	return &middlewares{
		jwtKey: key,
	}
}

func (mdd *middlewares) SetCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Authorization, origin, Content-Type, accept")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Content-Type", "application/json")

		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusOK)
		}
	}
}

func (mdd *middlewares) AuthorizeRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		authValue := strings.Split(authHeader, " ")
		if len(authValue) != 2 {
			log.Error().Stack().Err(errors.New("invalid auth header")).Msg("")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(authValue[1], mdd.keyFunc)
		if err != nil {
			log.Error().Stack().Err(errors.New("token parse failed")).Msg("")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			log.Error().Stack().Err(errors.New("invalid token")).Msg("")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Warn().Msg("malformed token")
			return
		}

		username, exists := claims["iss"].(string)
		if !exists {
			log.Warn().Msg("unidentified user")
			return
		}

		ctx.Set("username", username)
	}
}

func (mdd *middlewares) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(mdd.jwtKey), nil
}

func (mdd *middlewares) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		clientIP := c.ClientIP()

		// Log incoming request
		log.Info().
			Str("method", method).
			Str("path", path).
			Str("client_ip", clientIP).
			Time("start", start).
			Msg("Request received")

		// Continue processing
		c.Next()

		userID, exists := c.Get("username")
		if !exists {
			userID = "unknown"
		}

		// Log response
		end := time.Now()
		latency := end.Sub(start)
		statusCode := c.Writer.Status()

		log.Info().
			Str("method", method).
			Str("path", path).
			Str("username", fmt.Sprint(userID)).
			Str("client_ip", clientIP).
			Int("status", statusCode).
			Time("start", start).
			Time("end", end).
			Dur("latency", latency).
			Msg("Request completed")
	}
}
