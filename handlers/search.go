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

func SearchHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "query parameter is required", http.StatusBadRequest)
		return
	}

	limit := 10
	offset := 0

	if l := r.URL.Query().Get("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil && val > 0 {
			limit = val
		}
	}
	if o := r.URL.Query().Get("offset"); o != "" {
		if val, err := strconv.Atoi(o); err == nil && val >= 0 {
			offset = val
		}
	}

	ctx := context.Background()

	conn, err := db.Pool.Acquire(ctx)
	if err != nil {
		http.Error(w, "DB acquire error", http.StatusInternalServerError)
		return
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, queries.SearchSQL, query, limit, offset)
	if err != nil {
		http.Error(w, fmt.Sprintf("Query error: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []models.Film

	for rows.Next() {
		var f models.Film
		if err := rows.Scan(&f.ID, &f.Title, &f.Description, &f.Actors, &f.Genre, &f.ReleaseYear); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		results = append(results, f)
	}

	if len(results) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
