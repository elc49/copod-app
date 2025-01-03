# Run server/api tests
test-copod:
	cd Server && go test ./test && cd ../SmartContract && npm run test
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
	cd SmartContract && rm -rf ignition/deployments && npx hardhat ignition deploy ./ignition/modules/Registry.ts --network sepolia && solc --base-path . --include-path node_modules --abi --bin --overwrite contracts/Registry.sol -o ../Server/abi && cd ../Server && abigen --abi abi/Registry.abi --bin abi/Registry.bin --pkg registry --out contract/registry.go
