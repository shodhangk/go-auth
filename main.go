package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/shodhangk/go-auth/models"
	"github.com/shodhangk/go-auth/routes"

	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(e)

	port := os.Getenv("PORT")
	models.InitDB()
	// Handle routes
	http.Handle("/", routes.Handlers())

	// serve
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
