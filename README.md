go build -o out && ./out

psql "postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable"

goose postgres "postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable" up

goose postgres "postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable" down
