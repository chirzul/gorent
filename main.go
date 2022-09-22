package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chirzul/gorent/src/routers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	mainRoute, err := routers.New()
	if err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("APP_PORT")

	fmt.Println("service run on port", PORT)
	http.ListenAndServe(PORT, mainRoute)
}
