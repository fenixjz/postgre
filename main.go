package main

import (
	"fmt"
	"log"
	"net/http"

	"postgre/config"
	"postgre/db"
	"postgre/handlers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Initialize DB connection
	db.Connect()
	defer db.Close()

	// Initialize router
	router := httprouter.New()
	router.GET("/search", handlers.SearchHandler)

	fmt.Println("✅ Server is running at " + config.HTTP_URL + ":" + config.HTTP_PORT)

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":"+config.HTTP_PORT, router))
}
