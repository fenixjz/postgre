package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"postgre/db"
	"postgre/models"
	"postgre/queries"

	"github.com/julienschmidt/httprouter"
)

// SearchHandler performs full-text search on the films table.
//
// It reads the query string and optional limit/offset parameters.
// It returns a JSON array of Film objects or appropriate HTTP status.
func SearchHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "query parameter is required", http.StatusBadRequest)
		return
	}

	limit := 10
	offset := 0

	// Parse and validate optional limit parameter
	if l := r.URL.Query().Get("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil && val > 0 {
			limit = val
		}
	}

	// Parse and validate optional offset parameter
	if o := r.URL.Query().Get("offset"); o != "" {
		if val, err := strconv.Atoi(o); err == nil && val >= 0 {
			offset = val
		}
	}

	// Execute search query from embedded SQL file
	rows, err := db.Conn.Query(context.Background(), queries.SearchSQL, query, limit, offset)
	if err != nil {
		http.Error(w, fmt.Sprintf("Query error: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []models.Film

	// Map rows to Film structs
	for rows.Next() {
		var f models.Film
		if err := rows.Scan(&f.ID, &f.Title, &f.Description, &f.Actors, &f.Genre, &f.ReleaseYear); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		results = append(results, f)
	}

	// If no results, return 204 No Content
	if len(results) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Respond with JSON result
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
