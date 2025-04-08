// Package queries sadrži SQL upite koji se koriste u aplikaciji.
// Ovaj modul omogućava enkapsulaciju SQL koda tako da se upiti učitavaju
// direktno iz eksternih datoteka, što olakšava održavanje i modifikaciju upita.
package queries

// Import _ "embed" je potrebna linija da bi se koristile Go-ove direktive za embedovanje,
// tj. da bi se sadržaj SQL fajlova uvrstio u binarni fajl prilikom kompajliranja.
import _ "embed"

// SearchSQL sadrži sadržaj SQL upita koji se koristi za pretragu filmova.
// Direktiva //go:embed omogućava da se sadržaj datoteke search_films.sql uvrsti u promenljivu
// prilikom kompajliranja, tako da upit nije potrebno posebno čuvati kao eksterni fajl.
//
//go:embed search_films.sql
var SearchSQL string
