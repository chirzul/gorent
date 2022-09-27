package middlewares

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/chirzul/gorent/src/libs"
)

func Upload(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, handlerFile, err := r.FormFile("image")
		if err != nil {
			libs.GetResponse(err.Error(), 400, true)
			return
		}
		defer file.Close()

		fileName := time.Now().Format(time.UnixDate) + handlerFile.Filename
		fileDestination, err := os.Create("uploads/" + fileName)
		if err != nil {
			libs.GetResponse(err.Error(), 400, true)
			return
		}

		_, err = io.Copy(fileDestination, file)
		if err != nil {
			libs.GetResponse(err.Error(), 400, true)
			return
		}

		ctx := context.WithValue(r.Context(), "imageName", "uploads/"+fileName)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
