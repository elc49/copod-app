package contracts

import (
	"github.com/elc49/copod/config"
	"github.com/elc49/copod/contracts/land"
	"github.com/elc49/copod/contracts/registry"
	"github.com/elc49/copod/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

const (
	ZERO_ADDRESS = "0x0000000000000000000000000000000000000000"
)

var (
	eth EthBackend
)

type EthBackend interface {
	GetRegistryContract() *registry.Registry
	GetLandTitleDetails(string) (*land.LandDetails, error)
}

type ethBackend struct {
	client           *ethclient.Client
	registryContract *registry.Registry
	log              *logrus.Logger
}

func NewEthBackend() {
	log := logger.GetLogger()
	// Ethereum node connection
	conn, err := ethclient.Dial(config.C.Ethereum.InfuraApi)
	if err != nil {
		log.WithError(err).Fatalln("registry:Failed to connect to Ethereum client")
	} else {
		log.Infoln("registry:Connected to Ethereum client")
	}

	// Registry contract instance
	r, err := registry.NewRegistry(common.HexToAddress(config.C.Ethereum.RegistryContractAddress), conn)
	if err != nil {
		log.WithError(err).Fatalln("registry:Failed to instantiate a Registry contract")
	} else {
		log.Infoln("registry:Instantiated Registry smart contract")
	}

	eth = &ethBackend{conn, r, logger.GetLogger()}
}

func GetEthBackend() EthBackend {
	return eth
}

func (e *ethBackend) GetRegistryContract() *registry.Registry {
	return e.registryContract
}

func (e *ethBackend) GetLandTitleDetails(titleNo string) (*land.LandDetails, error) {
	contractAddress, err := e.registryContract.GetLandERC721Contract(nil, titleNo)
	if err != nil {
		e.log.WithError(err).WithFields(logrus.Fields{"title_no": titleNo}).Errorf("contract: GetLandERC721Contract: GetLandTitleDetails")
		return nil, err
	}

	landContract, err := land.NewLand(common.HexToAddress(contractAddress.String()), e.client)
	if err != nil {
		e.log.WithError(err).WithFields(logrus.Fields{"contract_address": contractAddress}).Errorf("contract: NewLand: GetLandTitleDetails")
		return nil, err
	}

	l, err := landContract.GetLand(nil)
	if err != nil {
		e.log.WithError(err).Errorf("contract: GetLand: GetLandTitleDetails")
		return nil, err
	}

	return &l, nil
}
