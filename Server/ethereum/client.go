package ethereum

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/elc49/copod/config"
	"github.com/elc49/copod/ethereum/land"
	"github.com/elc49/copod/ethereum/registry"
	"github.com/elc49/copod/logger"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	RegisterLand(context.Context, LandDetails) error
}

type LandDetails struct {
	TitleNo          string
	Symbol           string
	Owner            common.Address
	Size             *big.Int
	RegistrationDate *big.Int
}

type ethBackend struct {
	client           *ethclient.Client
	registryContract *registry.Registry
	log              *logrus.Logger
	signingAccount   *ecdsa.PrivateKey
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

	// Signing account
	privateKey, err := crypto.HexToECDSA(config.C.Ethereum.SigningAccountKey)
	if err != nil {
		log.WithError(err).Fatalln("ethereum: crypto.HexToECDSA: RegisterLand")
	}

	// Registry contract instance
	r, err := registry.NewRegistry(common.HexToAddress(config.C.Ethereum.RegistryContractAddress), conn)
	if err != nil {
		log.WithError(err).Fatalln("registry:Failed to instantiate a Registry contract")
	} else {
		log.Infoln("registry:Instantiated Registry smart contract")
	}

	eth = &ethBackend{conn, r, logger.GetLogger(), privateKey}
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

// signingPublicKey get public key from private key
func (e *ethBackend) signingPublicKey() common.Address {
	publicKey := e.signingAccount.Public()
	pkEcdsa, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		e.log.Fatalln("error casting public key to ecdsa")
	}

	return crypto.PubkeyToAddress(*pkEcdsa)
}

// getChainID for the current chain
func (e *ethBackend) getChainID(ctx context.Context) (*big.Int, error) {
	chainId, err := e.client.ChainID(ctx)
	if err != nil {
		e.log.WithError(err).Errorf("ethereum: ethclient.ChainID: getChainID")
		return nil, err
	}

	return chainId, nil
}

// Register land using registry contract
func (e *ethBackend) RegisterLand(ctx context.Context, l LandDetails) error {
	c, err := e.getChainID(ctx)
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(e.signingAccount, c)
	if err != nil {
		e.log.WithError(err).Errorf("ethereum: bind.NewKeyedTransactorWithChainID: RegisterLand")
		return err
	}

	tx, err := e.registryContract.Register(auth, l.TitleNo, l.Symbol, l.Owner, l.Size, l.RegistrationDate)
	if err != nil {
		e.log.WithError(err).WithFields(logrus.Fields{"land": l}).Errorf("registry: e.registryContract.Register: RegisterLand")
		return err
	}
	e.log.WithFields(logrus.Fields{"hash": tx.Hash()}).Infoln("success: register land")

	return nil
}
