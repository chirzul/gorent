package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	claim := NewToken("chirzul", "user")
	assert.IsType(t, &claims{}, claim)
	assert.Equal(t, "chirzul", claim.Username, "username must be chirzul")
	assert.Equal(t, "user", claim.Role, "role must be user")
}

func TestCreateToken(t *testing.T) {
	claim := NewToken("chirzul", "user")
	_, err := claim.Create()

	assert.Nil(t, err, "error must be nil")
}
