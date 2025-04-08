/*
 * main.go
 *
 * Glavna ulazna tačka za API server za pretragu filmova.
 * Ovaj fajl postavlja HTTP server, konfiguriše rute i povezuje
 * se sa PostgreSQL bazom podataka.
 */

package main

import (
	"fmt"
	"log"
	"net/http"

	"postgre/config"
	"postgre/db"
	"postgre/handlers"

	"github.com/julienschmidt/httprouter" // Router biblioteka za HTTP zahteve
)

func main() {
	// Povezivanje sa bazom podataka
	db.Connect()
	// Osiguravamo da će konekcija biti zatvorena kada program završi
	defer db.Close()

	// Inicijalizacija router-a
	router := httprouter.New()

	// Registrovanje endpoint-a za pretragu filmova
	router.GET("/search", handlers.SearchHandler) // Endpoint za pretragu po ključnim rečima

	// Štampanje informacije o pokretanju servera
	fmt.Println("✅ Server is running at " + config.HTTP_URL + ":" + config.HTTP_PORT)

	// Pokretanje HTTP servera (blokira izvršavanje)
	// U slučaju greške, log.Fatal će zapisati grešku i zaustaviti program
	log.Fatal(http.ListenAndServe(":"+config.HTTP_PORT, router))
}
