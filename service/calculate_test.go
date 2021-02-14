package service

import (
	"../models"
	"strconv"
	"strings"
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
		start = strings.Replace(start, "0x", "", -1)
		//fmt.Println("1", start)

		result, err := strconv.ParseUint(start, 16, 64)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		//fmt.Println("2", result)

		amountTransactions += float64(result)
	}

	amountStr := strconv.Itoa(int(amountTransactions))
	//fmt.Println("3", amountStr)

	t.Log("count:", countTransactions, "amount", amountStr)
}
