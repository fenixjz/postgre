/*
 * config.go
 *
 * Konfiguracioni fajl koji sadrži konstante potrebne za rad aplikacije.
 * Definiše parametre za HTTP server i konekciju sa bazom podataka.
 *
 * NAPOMENA: U produkcijskom okruženju, osetljive informacije kao što su
 * kredencijali za bazu podataka ne bi trebalo da budu hardkodirani,
 * već bi trebalo da se učitavaju iz okruženja (env varijabli) ili
 * iz konfiguracione datoteke koja nije u repozitorijumu.
 */

package config

// Konstante za konfiguraciju aplikacije
const (
	// HTTP konfiguracija
	HTTP_URL  = "http://localhost" // Bazni URL za server
	HTTP_PORT = "8200"             // Port na kojem će server osluškivati zahteve

	// Parametri za konekciju sa bazom podataka
	DB_USER     = "postgres"     // Korisničko ime za pristup bazi
	DB_PASSWORD = "3011972Fenix" // Lozinka za pristup bazi (bolje koristiti env varijable)
	DB_HOST     = "localhost"    // Adresa servera baze podataka
	DB_PORT     = "5432"         // Port za PostgreSQL servis
	DB_NAME     = "filmoteka"    // Naziv baze podataka koja se koristi
)
