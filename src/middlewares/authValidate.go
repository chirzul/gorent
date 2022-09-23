package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/chirzul/gorent/src/libs"
)

func CheckAuth(next http.HandlerFunc, roles []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		headerToken := r.Header.Get("Authorization")
		if !strings.Contains(headerToken, "Bearer") {
			libs.GetResponse("invalid header type", 401, true).Send(w)
			return
		}

		token := strings.Replace(headerToken, "Bearer ", "", -1)

		checkToken, err := libs.CheckToken(token)
		if err != nil {
			libs.GetResponse(err.Error(), 401, true).Send(w)
			return
		}

		ctx := context.WithValue(r.Context(), "username", checkToken.Username)

		var isAllowed bool
		for _, role := range roles {
			if role == checkToken.Role {
				isAllowed = true
			}
		}

		if !isAllowed {
			libs.GetResponse("you don't have access", 403, true).Send(w)
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
