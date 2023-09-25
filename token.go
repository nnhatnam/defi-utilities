package defi

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

type Token struct {
	chainID  ChainID
	address  common.Address
	decimals int64
	symbol   string
	name     string

	// bypassChecksum If true it only checks for length === 42, startsWith 0x and contains only hex characters
	bypassChecksum bool
}

// ensure Token conforms to BaseCurrency interface
//var _ Currency = &Token{}

func NewToken(chainID ChainID, address string, decimals int64, symbol string, name string, bypassChecksum bool) *Token {
	var hexAddress common.Address
	if address != "" {
		hexAddress = common.HexToAddress(address)
	}
	return &Token{
		chainID:        chainID,
		address:        hexAddress,
		decimals:       decimals,
		symbol:         symbol,
		name:           name,
		bypassChecksum: bypassChecksum,
	}
}

func (currency *Token) IsNative() bool {
	return false
}

func (currency *Token) IsToken() bool {
	return true
}

func (currency *Token) ChainID() ChainID {
	return currency.chainID
}

func (currency *Token) Address() common.Address {
	return currency.address
}

func (currency *Token) Decimals() int64 {
	return currency.decimals
}

func (currency *Token) Symbol() string {
	return currency.symbol
}

func (currency *Token) Name() string {
	return currency.name
}

func (currency *Token) Equals(other *Token) bool {
	// Implement the comparison logic here
	return other.IsToken() && currency.chainID == other.ChainID() && strings.ToLower(currency.address.Hex()) == strings.ToLower(other.Address().Hex())
}

func (currency *Token) GetWrapped() *Token {
	return currency
}

func (currency *Token) SortsBefore(other *Token) (bool, error) {
	if currency.ChainID() != other.ChainID() {
		return false, errors.New("chainId mismatch")
	}

	if currency.Address() == other.Address() {
		return false, errors.New("identical addresses")
	}

	return strings.ToLower(currency.Address().String()) < strings.ToLower(other.Address().String()), nil
}
