package service

import (
	"../models"
	"fmt"
	"math/big"
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

		n := new(big.Float)
		fl, _ := n.SetString(start)
		s := fmt.Sprintf("%e", fl)

		fl2, err := strconv.ParseFloat(s, 64)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		amountTransactions += fl2
	}

	s := fmt.Sprintf("%e", amountTransactions)
	amount := strings.Split(s, "e")

	if len(amount) == 0 {
		return
	}

	amountTransactionsStr := amount[0]

	t.Log("count:", countTransactions, "amount", amountTransactionsStr)
}
