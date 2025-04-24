package main

import (
	"github.com/deyvissonbrenoveras/aws-scalable-fintech/transactions-processor/processor"
	"github.com/deyvissonbrenoveras/aws-scalable-fintech/transactions-processor/transactions"
)

func main() {

	newTransaction := transactions.Transaction{
		TransactionID: "facc880a-bb63-4133-bc08-ac8f74492491",
		Amount:        1000,
		UserId:        "c52f3e4b0-8c1f-4a2b-9d3f-5e6c7d8e9f0",
		AccountID:     "account123",
		Type:          "TRANSFER",
		Method:        "PIX",
		Destination:   "pixkey123",
	}
	processor.ProcessTransaction(newTransaction)

	
}