package main

import (
	"fmt"
	"sync"

	"github.com/deyvissonbrenoveras/aws-scalable-fintech/transactions-processor/database"
	"github.com/deyvissonbrenoveras/aws-scalable-fintech/transactions-processor/processor"
	"github.com/deyvissonbrenoveras/aws-scalable-fintech/transactions-processor/transactions"
	"github.com/google/uuid"
)

func main() {

	var wg sync.WaitGroup

	for i :=range 10000 {
		wg.Add(1)
		newTransaction := transactions.Transaction{
			Id:            uuid.New().String(),
			Amount:        1000,
			UserId:        uuid.New().String(),
			AccountID:     "account123",
			Type:          "TRANSFER",
			Method:        "PIX",
			Destination:   "pixkey123",
			Status: "PENDING",
		}		
		
		fmt.Printf("Start processing transaction %d\n", i)
		go paralelWorker(&newTransaction, &wg)
	}
	wg.Wait() 
	fmt.Printf("Itens salvos na base %d", len(database.GetAllTransactions()))
	
}

func paralelWorker(transaction *transactions.Transaction, wg *sync.WaitGroup){	
	defer wg.Done()
	processor.ProcessTransaction(transaction)
}