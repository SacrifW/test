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
		log.Println("api.go -> CalculateTotal() -> error:", err)
		return
	}

	countTransactions, amountTransactions := service.CalculateCountAndAmountOfTransactions(blockNumber)

	c.JSON(200, models.Obj{"transactions": countTransactions, "amount": amountTransactions})
}