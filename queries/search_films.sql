SELECT id, title, description, actors, genre, release_year
FROM films
WHERE search_vector @@ plainto_tsquery('english'
    , $1)
    LIMIT $2
OFFSET $3;