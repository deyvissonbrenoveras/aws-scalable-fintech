package transactions

import (
	"errors"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID string `json:"transactionId"`
	Amount        int    `json:"amount"`
	UserId        string `json:"userId"`
	AccountID     string `json:"accountId"`
	Type          string `json:"type"`
	Method        string `json:"method"`
	Destination   string `json:"destination"`
}

const (
	TransactionTypeDeposit = "DEPOSIT"
	TransactionTypeWithdrawal = "WITHDRAWAL"
	TransactionTypeTransfer = "TRANSFER"
	TransactionTypePayment = "PAYMENT"	
	TransactionMethodPix = "PIX"
	TransactionMethodTed = "TED"
	TransactionMethodCashDeposit = "CASH_DEPOSIT"
	TransactionMethodCashWithdrawal = "CASH_WITHDRAWAL"
)

func (t *Transaction) Validate() error {
	var errorList []error
	
	if _, err := uuid.Parse(t.TransactionID); err != nil {
		errorList = append(errorList, errors.New("TransactionID is not a valid UUID"))
	}

	if t.Amount <= 0 {
		errorList = append(errorList, errors.New("amount must be greater than 0"))
	}

	if t.UserId == "" {
		errorList = append(errorList, errors.New("UserId cannot be empty"))
	}

	if t.AccountID == "" {
		errorList = append(errorList, errors.New("AccountID cannot be empty"))
	}

	if t.Type != TransactionTypeDeposit && t.Type != TransactionTypeWithdrawal && t.Type != TransactionTypeTransfer && t.Type != TransactionTypePayment {
		errorList = append(errorList, errors.New("invalid transaction type"))
	}

	if t.Method != TransactionMethodPix && t.Method != TransactionMethodTed && t.Method != TransactionMethodCashDeposit && t.Method != TransactionMethodCashWithdrawal {
		errorList = append(errorList, errors.New("invalid transaction method"))
	}

	if t.Destination == "" {
		errorList = append(errorList, errors.New("Destination cannot be empty"))
	}

	if len(errorList) != 0 {
		return errors.Join(errorList...)
	}
	
	return nil

}