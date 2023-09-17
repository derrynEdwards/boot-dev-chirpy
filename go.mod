module github.com/derrynEdwards/boot-dev-chirpy

go 1.21.1

replace github.com/derrynEdwards/boot-dev-chirpy/internal/database v0.0.0 => ./internal/database

replace github.com/derrynEdwards/boot-dev-chirpy/internal/auth v0.0.0 => ./internal/auth

require (
	github.com/go-chi/chi/v5 v5.0.10 // indirect
	github.com/golang-jwt/jwt/v5 v5.0.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/crypto v0.13.0 // indirect
)
