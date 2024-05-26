package infra

import (
	"time"

	"github.com/google/uuid"
)

type utils struct{}

func NewUtils() *utils {
	return &utils{}
}

func (ut *utils) GenerateID() string {
	return uuid.New().String()
}

func (ut *utils) GetCurrentTime() string {
	return time.Now().UTC().Format(time.RFC3339)
}
