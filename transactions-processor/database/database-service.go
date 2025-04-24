package database

import (
	"fmt"

	"github.com/deyvissonbrenoveras/aws-scalable-fintech/transactions-processor/transactions"
)

var transactionsDB []*transactions.Transaction

func GetTransactionById(transactionId string) *transactions.Transaction {
	for _, transaction := range transactionsDB {
		if transaction.Id == transactionId {
			return transaction
		}
	}
	return nil
}

func SaveTransaction(transaction *transactions.Transaction) error {
	fmt.Printf("Saving transaction to database with id %+v\n", transaction.Id)
	transactionsDB = append(transactionsDB, transaction)
	return nil
}

func GetAllTransactions() []*transactions.Transaction{
	return transactionsDB
}