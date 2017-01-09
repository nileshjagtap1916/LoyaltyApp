/*
Copyright Capgemini India. 2016 All Rights Reserved.
*/

package main

import (
	"errors"
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Loyalty Chaincode implementation
type LoyaltyChaincode struct {
}

func (t *LoyaltyChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	var err error
	// Initialize the chaincode
	
	fmt.Printf("Deployment of Loyalty is completed\n")

	var emptyMerchantTxs []MerchantTransactions
	jsonAsBytes, _ := json.Marshal(emptyMerchantTxs)
	err = stub.PutState(merchantIndexTxStr, jsonAsBytes)
	if err != nil {
		return nil, err
	}
	
	var emptyConsumerTxs []ConsumerTransactions
	jsonAsBytes, _ = json.Marshal(emptyConsumerTxs)
	err = stub.PutState(consumerIndexTxStr, jsonAsBytes)
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}

// Transaction makes payment of X units from A to B
func (t *LoyaltyChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == registerMerchantTxn {		
		return RegisterMerchant(stub, args)
	} else if function == registerConsumerTxn {
		return RegisterConsumer(stub, args)
	} else if function == addTransaction {
		return AddTransaction(stub, args)
	}
	return nil, nil
}



// Query callback representing the query of a chaincode
func (t *LoyaltyChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	var UserId string // Entities
	var err error
	var resAsBytes []byte

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	UserId = args[0]
	
	if function == "MERCHANT_TX_HISTORY" {		
		resAsBytes, err = GetMerchantTxHistory(stub, UserId)
	} else if function == "CONSUMER_TX_HISTORY" {
		resAsBytes, err = GetConsumerTxHistory(stub, UserId)
	}

	fmt.Printf("Query Response:%s\n", resAsBytes)
	
	if err != nil {
		return nil, err
	}
	
	return resAsBytes, nil
}



func main() {
	err := shim.Start(new(LoyaltyChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
