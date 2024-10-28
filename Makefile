# Run smart contract tests
tests:
	cd SmartContract && npm run test
# Tidy server
tidy-server:
	cd Server && go mod tidy
