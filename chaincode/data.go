package main

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Transaction struct{
	TransactionId string `json:"TXN_UUID"`
	MerchantID string `json:"PROGRAM_PARTNER_ID"`
	ConsumerID string `json:"TX_CONSUMER_ID"`
	TxnTokenName string `json:"TXN_PROGRAM_PARTNER_TOKEN_NAME"`
	TxnTokenQty int `json:"TXN_PROGRAM_PARTNER_TOKEN_QTY"`
	TxnDescription string `json:"TXN_DESC"`
	TxnType string `json:"TXN_TYPE"`
	TxTimestamp *timestamp.Timestamp `json:"TXN_TIMESTAMP"`
}

type MerchantTransactions struct{
	Name string `json:"name"`
	Transactions []Transaction `json:"transactions"`
}

type ConsumerTransactions struct{
	Name string `json:"name"`	
	Transactions []Transaction `json:"transactions"`
}

var merchantIndexTxStr = "_merchantIndexTxStr"
var consumerIndexTxStr = "_consumerIndexTxStr"
var registerMerchantTxn = "REG_MERCHANT"
var registerConsumerTxn = "REG_CONSUMER"
var addTransaction = "ADD_TX"