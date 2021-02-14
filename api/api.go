package api

import (
	"../models"
	"../service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
)

func CalculateTotal(c *gin.Context) {

	BlockNumberParam := c.Param("blockNumber")

	blockNumber, err := strconv.ParseUint(BlockNumberParam, 10, 64)
	if err != nil {
		log.Println("api.go -> CalculateTotal() -> ParseUint() -> error:", err)
		return
	}
	blockNumberFl := int(blockNumber)

	countTransactions, amountTransactions, err := service.CalculateCountAndAmountOfTransactions(blockNumberFl)
	if err != nil {
		log.Println("api.go -> CalculateTotal() -> CalculateCountAndAmountOfTransactions() -> error:", err)
		return
	}

	s := fmt.Sprintf("%e", amountTransactions)
	amount := strings.Split(s, "e")

	if len(amount) == 0 {
		return
	}
	amountTransactionsStr := amount[0]

	c.JSON(200, models.Obj{"transactions": countTransactions, "amount": amountTransactionsStr})
}
