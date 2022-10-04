package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	result, err := HashPassword(password)

	assert.True(t, result != password, "Password must be hashed")
	assert.Nil(t, err, "Error must be nil")
}

func TestCheckPassword(t *testing.T) {
	inputPass := "user"
	dbPass := "$2a$10$3EIKZ6FemRn2cblEioUUKuFDkEKDxd3RqkIWxXev7PEpLdoLaDwg2"
	assert.True(t, CheckPassword(dbPass, inputPass), "Result must be true")
	assert.False(t, CheckPassword(dbPass, "falsepassword"), "Result must be false")
}
