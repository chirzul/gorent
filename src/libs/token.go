package libs

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySecret = []byte(os.Getenv("JWT_KEYS"))

type claims struct {
	Username string
	Role     string
	jwt.StandardClaims
}

func NewToken(username, role string) *claims {
	return &claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
}

func (c *claims) Create() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return tokens.SignedString(mySecret)
}

func CheckToken(token string) (*claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims := tokens.Claims.(*claims)
	return claims, nil
}
