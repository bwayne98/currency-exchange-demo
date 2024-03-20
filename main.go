package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExchangeRequest struct {
	Source string `json:"source" validate:"required"`
	Target string `json:"target" validate:"required"`
	Amount string `json:"amount" validate:"required"`
}

type ExchangeResponse struct {
	Msg    string `json:"msg"`
	Amount string `json:"amount"`
}

func main() {
	route := gin.Default()

	route.SetTrustedProxies([]string{})

	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	route.Run(":80")
}
