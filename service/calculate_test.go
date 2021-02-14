package service

import (
	"../models"
	"fmt"
	"strconv"
	"testing"
)

func TestCalculateTotal(t *testing.T) {
	firstBlock := models.Block{
		Result: models.Result{Transactions: []models.Transaction{{Value: "0x38d7ea4c68000"}, {Value: "0x4edec84a038000"}, {Value: "0x2ee5547f090000"}}},
	}

	block := firstBlock //mock

	var amountTransactions float64

	countTransactions := len(block.Result.Transactions)

	for v := range block.Result.Transactions {
		start := block.Result.Transactions[v].Value
		fmt.Println("454454", start)
		//result, err := strconv.ParseInt("0x1c8", 0, 64)
		result, err := strconv.ParseFloat(start, 32)
		fmt.Println("sfsfsefs", result)
		if err != nil {
			t.Errorf("err: %v", err)
		}
		amountTransactions += result
	}

	t.Log("count:", countTransactions, "amount", amountTransactions)
}