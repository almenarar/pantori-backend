package infra

import (
	core "pantori/internal/auth/core"

	"os"
	"strings"
)

type memory struct {
	prefix string
}

func NewMemory() *memory {
	return &memory{
		prefix: "PANTORI_USER",
	}
}

// simple temp solution for user management
// storing them in env vars
func (mm *memory) GetUser(user core.User) (core.User, error) {
	envVars := os.Environ()

	// Filter variables that start with the specified prefix
	filteredVars := make(map[string]string)
	for _, envVar := range envVars {
		pair := strings.SplitN(envVar, "=", 2)
		if strings.HasPrefix(pair[0], mm.prefix) {
			filteredVars[pair[0]] = pair[1]
		}
	}

	for key, value := range filteredVars {
		if user.Username == strings.Split(key, "_")[2] {
			user.ActualPassword = value
			return user, nil
		}
	}

	return core.User{}, nil
}
