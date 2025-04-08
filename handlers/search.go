package handlers

import (
	"context"       // Omogućava rad sa kontekstom (context) za kontrolu timeouta, cancel signala itd.
	"encoding/json" // Za enkodiranje i dekodiranje podataka u JSON formatu.
	"fmt"           // Omogućava formatiranje stringova i kreiranje poruka.
	"log"           // Paket za logovanje informacija, upozorenja i grešaka.
	"net/http"      // Pruža funkcionalnosti HTTP servera, uključujući rad sa zahtevima i odgovorima.
	"strconv"       // Omogućava konverziju između stringova i numeričkih tipova (npr. parsiranje "limit" i "offset" parametara).

	"postgre/db"      // Lokalni modul za konekciju i rad sa PostgreSQL bazom.
	"postgre/models"  // Lokalni modul koji sadrži definicije modela podataka (npr. struktura Film).
	"postgre/queries" // Lokalni modul sa SQL upitima koji se koriste za pretragu i dohvat podataka iz baze.

	"github.com/julienschmidt/httprouter" // Eksterni paket koji omogućava brzu i jednostavnu rutiranje HTTP zahteva.
)

/*
SearchHandler funkcija obrađuje HTTP GET zahteve za pretragu filmova.

Ona vrši sledeće korake:
1. Preuzima i validira query parametre iz URL-a, kao što su:
  - "query": termin za pretragu (obavezan)
  - "lang": jezik pretrage (podržani: "english", "serbian", "spanish")
  - "limit": maksimalan broj rezultata (podrazumevano 10)
  - "offset": pomeraj rezultata (podrazumevano 0)

2. Kreira kontekst (context) za upite.
3. Dobija konekciju iz Deadpool connection pool-a.
4. Izvršava SQL upit koristeći parametre pretrage.
5. Iterira kroz vraćene redove, mapira svaki red u strukturu models.Film.
6. Vraća odgovor u JSON formatu, ili odgovara sa statusom 204 ako nema rezultata.
*/
func SearchHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Preuzimanje obaveznog parametra "query" iz URL-a
	query := r.URL.Query().Get("query")
	if query == "" {
		// Ako query nije prosleđen, vraća se HTTP 400 Bad Request
		http.Error(w, "query parameter is required", http.StatusBadRequest)
		return
	}

	// Postavljanje podrazumevanih vrednosti
	lang := "english" // podrazumevani jezik pretrage
	limit := 10       // podrazumevani broj rezultata
	offset := 0       // podrazumevani offset

	// Provera opcionalnog parametra "lang" - podržani jezici su: english, serbian, spanish
	if ln := r.URL.Query().Get("lang"); ln != "" {
		// Mapiranje podržanih jezika za jednostavnu validaciju
		supportedLang := map[string]bool{
			"english": true,
			"serbian": true,
			"spanish": true,
		}

		// Ako je prosleđeni jezik podržan, koristi se taj jezik
		if supportedLang[ln] {
			lang = ln
		} else {
			// Ako je jezik nepoznat, vraća se HTTP 400 Bad Request sa odgovarajućom porukom
			http.Error(w, fmt.Sprintf("Unsupported language: %s. Supported languages are: english, serbian, spanish", ln), http.StatusBadRequest)
			return
		}
	}

	// Parsiranje i validacija opcionog parametra "limit"
	if l := r.URL.Query().Get("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil && val > 0 {
			limit = val
		}
	}
	// Parsiranje i validacija opcionog parametra "offset"
	if o := r.URL.Query().Get("offset"); o != "" {
		if val, err := strconv.Atoi(o); err == nil && val >= 0 {
			offset = val
		}
	}

	// Kreiranje konteksta za bazne operacije
	ctx := context.Background()

	// Dobijanje konekcije iz connection pool-a
	conn, err := db.Pool.Acquire(ctx)
	if err != nil {
		// Ako dođe do greške prilikom dobijanja konekcije iz pool-a,
		// vraća se HTTP 500 Internal Server Error
		http.Error(w, "DB acquire error", http.StatusInternalServerError)
		return
	}
	// Osigurava se oslobadjanje konekcije nakon završetka korišćenja
	defer conn.Release()

	// Izvršavanje SQL upita za pretragu filmova koristeći parametre:
	// lang, query, limit i offset
	rows, err := conn.Query(ctx, queries.SearchSQL, lang, query, limit, offset)
	if err != nil {
		// Ako dođe do greške prilikom izvršavanja upita, vraća se HTTP 500 sa detaljima greške
		http.Error(w, fmt.Sprintf("Query error: %v", err), http.StatusInternalServerError)
		return
	}
	// Osigurava se zatvaranje rows objekta kada se završi sa iteracijom
	defer rows.Close()

	// Kreiranje niza rezultata tipa models.Film
	var results []models.Film

	// Iteracija kroz redove rezultata upita
	for rows.Next() {
		var f models.Film
		// Skidanje podataka iz svakog reda i mapiranje u strukturu f
		if err := rows.Scan(&f.ID, &f.Title, &f.Description, &f.Actors, &f.Genre, &f.ReleaseYear); err != nil {
			// Ako dođe do greške pri mapiranju, greška se loguje i prelazi se na naredni red
			log.Printf("Error scanning row: %v", err)
			continue
		}
		// Dodavanje uspešno mapiranog filma u listu rezultata
		results = append(results, f)
	}

	// Ako nema rezultata, vraća se HTTP 204 No Content
	if len(results) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Postavljanje odgovora sa Content-Type zaglavljem
	w.Header().Set("Content-Type", "application/json")
	// Enkodiranje rezultata u JSON format i slanje odgovora
	if err := json.NewEncoder(w).Encode(results); err != nil {
		// Ako dođe do greške pri enkodiranju u JSON, vraća se HTTP 500 Internal Server Error
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
