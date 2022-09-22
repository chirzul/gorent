package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/chirzul/gorent/src/libs"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
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
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
