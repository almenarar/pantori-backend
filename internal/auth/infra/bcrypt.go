package infra

import (
	core "pantori/internal/auth/core"

	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type cryptography struct {
	jwtKey string
}

func NewCryptography(jwtKey string) *cryptography {
	return &cryptography{
		jwtKey: jwtKey,
	}
}

func (c *cryptography) EncryptPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (c *cryptography) CheckPassword(stored, given string) error {
	err := bcrypt.CompareHashAndPassword([]byte(stored), []byte(given))
	if err != nil {
		return err
	}
	return nil
}

func (c *cryptography) GenerateToken(user core.User) (string, error) {
	claims := jwt.MapClaims{
		"iss":       "pantori",
		"id":        uuid.New().String(),
		"sub":       user.Username,
		"exp":       time.Now().Add(time.Hour).Unix(),
		"iat":       time.Now().Unix(),
		"nbf":       time.Now().Unix(),
		"workspace": user.Workspace,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(c.jwtKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
