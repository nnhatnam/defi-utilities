package defi

import "math/big"

//go:generate stringer -type=ChainID -output=chain_string.go -linecomment
type ChainID uint64

func (c ChainID) Big() *big.Int {
	return big.NewInt(int64(c))
}

const (
	MAINNET         ChainID = 1        // Mainnet
	GOERLI          ChainID = 5        // Goerli
	SEPOLIA         ChainID = 11155111 // Sepolia
	OPTIMISM        ChainID = 10       // Optimism
	OPTIMISM_GOERLI ChainID = 420      // Optimism Goerli
	ARBITRUM_ONE    ChainID = 42161    // Arbitrum One
	ARBITRUM_GOERLI ChainID = 421611   // Arbitrum Goerli
	POLYGON         ChainID = 137      // Polygon
	POLYGON_MUMBAI  ChainID = 80001    // Polygon Mumbai
	CELO            ChainID = 42220    // Celo
	CELO_ALFAJORES  ChainID = 44787    // Celo Alfajores
	GNOSIS          ChainID = 60       // Gnosis
	MOONBEAM        ChainID = 1284     // Moonbeam
	BNB             ChainID = 56       // Binance Smart Chain
	AVALANCHE       ChainID = 43114    // Avalanche
	BASE_GOERLI     ChainID = 84531    // Base Goerli
	BASE            ChainID = 8453     // Base
	KAVA_EVM        ChainID = 2222     // Kava
)
