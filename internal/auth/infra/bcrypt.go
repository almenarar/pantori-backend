package infra

import (
	core "pantori/internal/auth/core"

	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type cryptography struct {
	jwt_key string
}

func NewCryptography(key string) *cryptography {
	return &cryptography{
		jwt_key: key,
	}
}

func (c *cryptography) CheckPassword(stored, given string) error {
	if stored == given {
		return nil
	}
	return errors.New("oops")
}

func (c *cryptography) GenerateToken(user core.User) (string, error) {
	claims := jwt.StandardClaims{
		Id:        uuid.New().String(),
		Audience:  "foo",
		Issuer:    user.Username,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		NotBefore: time.Now().Unix(),
		Subject:   fmt.Sprintf("%s-api-access", "foo"),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(c.jwt_key))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
