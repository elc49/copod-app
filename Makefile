# Run smart contract tests
test-contracts:
	cd SmartContract && npm run test
# Run server/api tests
test-api:
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
# Web
web-app:
	cd web && npm run dev
# Deploy smart contract to optimisim sepolia testnet
op-sepolia-deploy:
	cd SmartContract && rm -rf ignition/deployments && npx hardhat ignition deploy ./ignition/modules/Registry.ts --network sepolia
