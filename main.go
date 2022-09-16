package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chirzul/gorent/src/routers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mainRoute, err := routers.New()
	if err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")

	fmt.Println("service run on port", PORT)
	http.ListenAndServe(PORT, mainRoute)
}
