package main

import (
	"encoding/json"
	"fmt"

	"github.com/BalusChen/IKHNAIE/types"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type CCTransaction struct{}

func main() {
	err := shim.Start(new(CCTransaction))
	if err != nil {
		panic(err)
	}
}

func (cc *CCTransaction) Name() string {
	return "CCTransaction"
}

func (cc *CCTransaction) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (cc *CCTransaction) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	switch fn {
	case "addTransaction":
		return addTransaction(stub, args)
	case "getTransactionHistory":
		return getTransactionHistory(stub, args)
	default:
		return shim.Error(fmt.Sprintf("unknown function %q", fn))
	}
}

func addTransaction(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 2 {
		return shim.Error(fmt.Sprintf("wrong number of args, expect: %d, got: %d", 2, len(args)))
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(fmt.Sprintf("put state failed, err: %v", err))
	}
	return shim.Success(nil)
}

func getTransactionHistory(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		return shim.Error(fmt.Sprintf("wrong number of args, expect: %d, got: %d", 1, len(args)))
	}

	resultIterator, err := stub.GetHistoryForKey(args[0])
	if err != nil {
		return shim.Error(fmt.Sprintf("get history for key %q failed, err: %v", args[0], err))
	}
	defer resultIterator.Close()

	history := make([]types.Transaction, 0)
	var transaction types.Transaction
	for resultIterator.HasNext() {
		resp, err := resultIterator.Next()
		if err != nil {
			return shim.Error(fmt.Sprintf("get transaction history for key %q failed, err: %v", args[0], err))
		}

		_ = json.Unmarshal(resp.Value, &transaction)
		history = append(history, transaction)
	}

	data, _ := json.Marshal(history)
	return shim.Success(data)
}
