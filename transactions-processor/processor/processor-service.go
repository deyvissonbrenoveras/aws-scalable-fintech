package processor

import (
	"errors"
	"fmt"

	"github.com/deyvissonbrenoveras/aws-scalable-fintech/transactions-processor/database"
	"github.com/deyvissonbrenoveras/aws-scalable-fintech/transactions-processor/transactions"
)


func ProcessTransaction(transaction transactions.Transaction) (string, error) {

	fmt.Printf("Processing transaction: %+v\n", transaction)
	
	if err := transaction.Validate(); err != nil {
		validationErrors := fmt.Errorf("Transaction validation failed: %v\n", err)
		fmt.Printf(validationErrors.Error())
		return "", errors.New("invalid transaction")
	}

	if err := database.SaveTransaction(transaction); err != nil {
		fmt.Printf("Failed to save transaction: %v\n", err)
		return "", errors.New("failed to save transaction")
	}
	
	return "Transaction processed successfully", nil
}