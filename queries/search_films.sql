SELECT id, title, description, actors, genre, release_year
FROM films
WHERE search_vector @@ plainto_tsquery($1
    , $2)
    LIMIT $3
OFFSET $4;