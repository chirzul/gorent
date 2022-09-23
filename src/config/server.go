package config

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/chirzul/gorent/src/routers"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start application",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {

		address := "0.0.0.0:8080"
		PORT := os.Getenv("APP_PORT")

		if PORT != "" {
			address = ":" + PORT
		}

		srv := &http.Server{
			Addr:         address,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Minute,
			Handler:      mainRoute,
		}

		fmt.Println("server running at port", PORT)
		srv.ListenAndServe()
		return nil

	} else {
		return err
	}
}
