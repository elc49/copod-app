package util

import (
	"reflect"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

// IsValidAddress validate eth address
func IsValidAddress(address interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := address.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// IsZeroAddress validate if zero eth address
func IsZeroAddress(a interface{}) bool {
	if !IsValidAddress(a) {
		return false
	}

	var address common.Address
	switch v := a.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(zeroAddressBytes, addressBytes)
}
