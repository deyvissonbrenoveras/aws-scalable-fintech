package processor

import (
	"errors"
	"fmt"
	"time"

	"github.com/deyvissonbrenoveras/aws-scalable-fintech/transactions-processor/database"
	"github.com/deyvissonbrenoveras/aws-scalable-fintech/transactions-processor/transactions"
)


func ProcessTransaction(newTransaction *transactions.Transaction) error {
	time.Sleep(10000000000)
	var err error

	fmt.Printf("Processing transaction: %+v\n", newTransaction)
	
	if err = newTransaction.Validate(); err != nil {
		validationErrors := fmt.Errorf("transaction validation failed: %w", err)
		fmt.Println(validationErrors.Error())
		return  errors.New("invalid transaction")
	}

	fmt.Printf("Checking transaction idempotence with id %s\n", newTransaction.Id)
	
	existingTransaction := database.GetTransactionById(newTransaction.Id)

	if existingTransaction != nil {
		err = fmt.Errorf("transaction with id %s was already processed", newTransaction.Id)
		fmt.Println(err)
		return err
	}


	if err := database.SaveTransaction(newTransaction); err != nil {
		fmt.Printf("Failed to save transaction: %v\n", err)
		return  errors.New("failed to save transaction")
	}
	
	return nil
}