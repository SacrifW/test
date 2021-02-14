package api

import (
	"../models"
	"../service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func CalculateTotal(c *gin.Context) {

	BlockNumberParam := c.Param("blockNumber")

	blockNumber, err := strconv.ParseUint(BlockNumberParam, 10, 64)
	if err != nil {
		log.Println("api.go -> CalculateTotal() -> ParseUint() -> error:", err)
		return
	}
	blockNumberFl := float32(blockNumber)

	countTransactions, amountTransactions, err := service.CalculateCountAndAmountOfTransactions(blockNumberFl)
	if err != nil {
		log.Println("api.go -> CalculateTotal() -> CalculateCountAndAmountOfTransactions() -> error:", err)
		return
	}

	c.JSON(200, models.Obj{"transactions": countTransactions, "amount": amountTransactions})
}