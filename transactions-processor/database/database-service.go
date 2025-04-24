package database

import (
	"fmt"

	"github.com/deyvissonbrenoveras/aws-scalable-fintech/transactions-processor/transactions"
)

var transactionsDB []transactions.Transaction

func SaveTransaction(transaction transactions.Transaction) error {
	fmt.Printf("Saving transaction to database: %+v\n", transaction)
	transactionsDB = append(transactionsDB, transaction)
	return nil
}