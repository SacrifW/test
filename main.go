package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"strconv"
)

type (
	obj map[string]interface{}

	Block struct {
		jsonrpc string
		id      string
		result  Result
	}

	Result struct {
		difficulty      string
		extraData       string
		gasLimit        string
		gasUsed         string
		hash            string
		logsBloom       string
		miner           string
		mixHash         string
		nonce           string
		number          string
		parentHash      string
		receiptsRoot    string
		sha3Uncles      string
		size            string
		stateRoot       string
		timestamp       string
		totalDifficulty string
		transactions    []Transaction
	}
	Transaction struct {
		blockHash        string
		blockNumber      string
		from             string
		gas              string
		gasPrice         string
		hash             string
		input            string
		nonce            string
		to               string
		transactionIndex string
		value            string
		v                string
		r                string
		s                string
	}

)

func main() {

	r := gin.Default()
	r.GET("/api/block/:blockNumber/total", CalculateTotal)

	r.Run()
}

//подсчет
func CalculateTotal(c *gin.Context) {

	blockNumber := c.Param("blockNumber")

	block := GetBlock(blockNumber)

	var amountTransactions int64

	countTransactions := len(block.result.transactions)

	for v := range block.result.transactions {
		start := block.result.transactions[v].value
		result, err := strconv.ParseInt(start, 0, 64)
		if err != nil {
			fmt.Println("err", err)
		}
		amountTransactions += result
	}

	c.JSON(200, obj{"transactions": countTransactions, "amount": amountTransactions})
}

//обращение в базу
func GetBlock(blockNumber string) Block {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		fmt.Println("main.go -> GetBlock() -> error-db connect", err)
	}
	defer session.Close()

	// получаем коллекцию products в базе данных productdb
	BlockCollection := session.DB("blockDB").C("blocks")

	var block Block
	search := obj{"blockNumber": blockNumber}

	if err := BlockCollection.Find(search).One(&block); err != nil {
		fmt.Println("main.go -> GetBlock() -> error-db", err)
		return block
	}

	return block
}
