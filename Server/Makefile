# Tidy server
tidy:
	go mod tidy
# Make server
server:
	go run ./server.go
# Graphql
graphql:
	go generate ./...
# Make postgres
sqlc:
	sqlc generate
# Test
tests:
	go test ./test
