package middlewares

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/chirzul/gorent/src/libs"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func CloudinaryUpload(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(20 << 20); err != nil {
			libs.GetResponse(err.Error(), 400, true).Send(w)
			return
		}

		file, handlerFile, err := r.FormFile("image")
		defer file.Close()
		if err != nil {
			libs.GetResponse(err.Error(), 400, true).Send(w)
		}

		fileName := time.Now().Format(time.RFC822) + handlerFile.Filename

		ctx := context.Background()
		cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))

		resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "gorent-all", PublicID: fileName})
		if err != nil {
			libs.GetResponse(err.Error(), 400, true).Send(w)
		}

		cntx := context.WithValue(r.Context(), "imageName", resp.SecureURL)
		next.ServeHTTP(w, r.WithContext(cntx))
	}

}
