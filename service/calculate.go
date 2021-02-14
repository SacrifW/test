package service

import (
	"../db"
	"fmt"
	"log"
	"strconv"
	"strings"
)

 func CalculateCountAndAmountOfTransactions (blockNumber int) (int, float64, error) {
	 var amountTransactions float64
	 block, err := db.GetBlock(blockNumber)
	 if err != nil {
	 	log.Println("calculate.go -> CalculateCountAndAmountOfTransactions() -> GetBlock() -> error: ", err)
	 	return 0, 0, err
	 }

	 countTransactions := len(block.Result.Transactions)

	 for v := range block.Result.Transactions {
		 start := block.Result.Transactions[v].Value
		 start = strings.Replace(start, "0x", "", -1)
		 result, err := strconv.ParseUint(start, 16, 64)
		 if err != nil {
			 fmt.Println("calculate.go -> CalculateCountAndAmountOfTransactions() -> ParseFloat() -> error: ", err)
		 }
		 amountTransactions += float64(result)
	 }

	 return countTransactions, amountTransactions, nil
 }