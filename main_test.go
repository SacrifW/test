package main

import (
	"strconv"
	"testing"
)

func TestCalculateTotal(t *testing.T) {
	firstBlock := Block{
		result: Result{transactions: []Transaction{{value: "0x38d7ea4c68000"}, {value: "0x4edec84a038000"}, {value: "0x2ee5547f090000"}}},
	}

	block := firstBlock //mock

	var amountTransactions int64

	countTransactions := len(block.result.transactions)

	for v := range block.result.transactions {
		start := block.result.transactions[v].value
		result, err := strconv.ParseInt(start, 0, 64)
		if err != nil {
			t.Errorf("err: %v", err)
		}
		amountTransactions += result
	}

	t.Log("count:", countTransactions, "amount", amountTransactions)
}
