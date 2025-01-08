package contracts

import (
	"github.com/elc49/copod/config"
	"github.com/elc49/copod/contracts/registry"
	"github.com/elc49/copod/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	ZERO_ADDRESS = "0x0000000000000000000000000000000000000000"
)

var (
	client           *ethclient.Client
	registryContract *registry.Registry
)

type ethBackend struct{}

func NewEthBackend() {
	log := logger.GetLogger()
	// Ethereum node connection
	conn, err := ethclient.Dial(config.C.Ethereum.InfuraApi)
	if err != nil {
		log.WithError(err).Fatalln("registry:Failed to connect to Ethereum client")
	} else {
		log.Infoln("registry:Connected to Ethereum client")
	}
	client = conn

	// Registry contract instance
	r, err := registry.NewRegistry(common.HexToAddress(config.C.Ethereum.RegistryContractAddress), conn)
	if err != nil {
		log.WithError(err).Fatalln("registry:Failed to instantiate a Registry contract")
	} else {
		log.Infoln("registry:Instantiated Registry smart contract")
	}

	registryContract = r
}

func GetEthClient() *ethclient.Client {
	return client
}

func GetRegistryContract() *registry.Registry {
	return registryContract
}
