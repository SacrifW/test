package service

import (
	"../db"
	"fmt"
	"log"
	"strconv"
)

 func CalculateCountAndAmountOfTransactions (blockNumber float32) (int, float64, error) {
	 var amountTransactions float64
	 block, err := db.GetBlock(blockNumber)
	 if err != nil {
	 	log.Println("calculate.go -> CalculateCountAndAmountOfTransactions() -> GetBlock() -> error: ", err)
	 	return 0, 0, err
	 }

	 countTransactions := len(block.Result.Transactions)

	 for v := range block.Result.Transactions {
		 start := block.Result.Transactions[v].Value
		 result, err := strconv.ParseFloat(start, 64)
		 if err != nil {
			 fmt.Println("calculate.go -> CalculateCountAndAmountOfTransactions() -> ParseFloat() -> error: ", err)
		 }
		 amountTransactions += result
	 }

	 return countTransactions, amountTransactions, nil
 }