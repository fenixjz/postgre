package models

/*
Film predstavlja strukturu entiteta film. Ova struktura mapira sve bitne informacije
o filmu, koje se kasnije enkodiraju u JSON format ili koriste prilikom rada sa bazom podataka.
Tagovi (npr. `json:"id"`) omogućavaju automatsko mapiranje polja prilikom enkodiranja i dekodiranja
JSON objekata, kako bi nazivi polja u JSON-u bili prilagođeni (npr. "release_year" umesto "ReleaseYear").
*/
type Film struct {
	ID          int    `json:"id"`           // Jedinstveni identifikator filma.
	Title       string `json:"title"`        // Naslov filma.
	Description string `json:"description"`  // Opis filma, koji može sadržati sinopsis ili dodatne informacije.
	Actors      string `json:"actors"`       // Lista glavnih glumaca ili opis glumačke ekipe (može biti i niskonivska reprezentacija).
	Genre       string `json:"genre"`        // Žanr filma, npr. drama, komedija, triler itd.
	ReleaseYear int    `json:"release_year"` // Godina kada je film izašao u produkciju ili premijerno prikazan.
}
