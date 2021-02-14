package main

import (
	"./api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/block/:blockNumber/total", api.CalculateTotal)

	r.Run()
}
