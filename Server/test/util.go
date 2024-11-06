package test

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/elc49/copod/logger"
	"github.com/google/uuid"
)

func RandomStringByLength(length int) string {
	log := logger.GetLogger()
	b := ""
	id, err := uuid.NewUUID()
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i <= length; i++ {
		randInt, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			log.Fatalln(err)
			break
		}
		b += string(id.String()[randInt.Int64()])
	}

	return b
}

func RandomWalletAddress() string {
	return fmt.Sprintf("0x41eD3Ce6DC13fD4F67Eb715f5c3B105B%s", RandomStringByLength(8))
}

func RandomEmailAddress() string {
	return fmt.Sprintf("%s@em.com", RandomStringByLength(4))
}

func RandomGovtID() string {
	return RandomStringByLength(8)
}
