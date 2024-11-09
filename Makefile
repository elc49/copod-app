# Run smart contract tests
test-contracts:
	cd SmartContract && npm run test
# Run server/api tests
test-server:
	cd Server && go test ./test
# Tidy server
tidy:
	cd Server && go mod tidy
# Make server
api-server:
	cd Server && go run ./main.go
# Graphql
graphql:
	cd Server && go generate ./...
# Make postgres
sqlc:
	cd Server && sqlc generate
