/*
 * connection.go
 *
 * Modul za upravljanje konekcijom sa bazom podataka.
 * Sadrži funkcije za uspostavljanje i zatvaranje veze sa PostgreSQL bazom,
 * kao i globalnu promenljivu za pool konekcija koji se koristi u celoj aplikaciji.
 */

package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool" // Biblioteka za efikasno upravljanje PostgreSQL konekcijama
	"postgre/config"                  // Uvoz lokalnog konfiguracionog paketa
)

// Pool je globalna promenljiva koja sadrži konekcije sa bazom podataka
// Koristi se u celoj aplikaciji za izvršavanje upita
var Pool *pgxpool.Pool

// Connect uspostavlja konekciju sa PostgreSQL bazom podataka
// koristeći parametre iz konfiguracionog fajla
func Connect() {
	// Formiranje connection string-a za PostgreSQL
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.DB_USER,     // Korisničko ime
		config.DB_PASSWORD, // Lozinka
		config.DB_HOST,     // Host adresa
		config.DB_PORT,     // Port
		config.DB_NAME,     // Naziv baze
	)

	var err error
	// Kreiranje pool-a konekcija sa bazom
	Pool, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		// U slučaju greške, prekidamo izvršavanje programa
		log.Fatalf("Failed to create connection pool: %v", err)
	}
}

// Close zatvara pool konekcija i oslobađa resurse
// Poziva se kada se aplikacija završava
func Close() {
	if Pool != nil {
		Pool.Close()
		// Napomena: ovde bi bilo dobro vratiti grešku ako postoji
		// i obraditi je u main funkciji
	}
}
