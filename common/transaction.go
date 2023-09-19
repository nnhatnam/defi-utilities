package common

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func GetSender(transaction *types.Transaction, chainId *big.Int) (common.Address, error) {
	signer := types.LatestSignerForChainID(chainId)
	return signer.Sender(transaction)
}

func DecodeTransactionInputDataToMap(contractABI *abi.ABI, data []byte, inputsMap map[string]interface{}) (*abi.Method, error) {
	methodSigData := data[:4]
	inputsSigData := data[4:]
	method, err := contractABI.MethodById(methodSigData)
	if err != nil {
		return nil, err
	}
	//inputsMap := make(map[string]interface{})
	if err = method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		return nil, err
	}
	return method, nil
}
