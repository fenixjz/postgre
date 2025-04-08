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
	db.Connect()
	defer db.Close()

	router := httprouter.New()
	router.GET("/search", handlers.SearchHandler)

	fmt.Println("✅ Server is running at " + config.HTTP_URL + ":" + config.HTTP_PORT)

	log.Fatal(http.ListenAndServe(":"+config.HTTP_PORT, router))
}
