package service

import (
	"../db"
	"fmt"
	"log"
	"math/big"
	"strconv"
)

func CalculateCountAndAmountOfTransactions(blockNumber int) (int, float64, error) {
	var amountTransactions float64
	block, err := db.GetBlock(blockNumber)
	if err != nil {
		log.Println("calculate.go -> CalculateCountAndAmountOfTransactions() -> GetBlock() -> error: ", err)
		return 0, 0, err
	}

	countTransactions := len(block.Result.Transactions)

	for v := range block.Result.Transactions {
		start := block.Result.Transactions[v].Value

		n := new(big.Float)
		fl, _ := n.SetString(start)
		s := fmt.Sprintf("%e", fl)

		fl2, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Println("calculate.go -> CalculateCountAndAmountOfTransactions() -> ParseFloat() -> error: ", err)
		}

		amountTransactions += fl2
	}

	return countTransactions, amountTransactions, nil
}
