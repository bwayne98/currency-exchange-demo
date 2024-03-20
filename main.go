package main

import (
	"fmt"
	"net/http"

	"github.com/bwayne98/currency-exchange-demo/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	rateService := service.CurencyExchngeRateService{}

	exchangeRate, err := rateService.Fetch()

	if err != nil {
		fmt.Printf("%e", err)
		return
	}

	changeService := service.CurrencyExchangeService{}

	route.POST("exchange", func(c *gin.Context) {
		var request ExchangeRequest
		validate := validator.New()

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, ExchangeResponse{
				Msg: err.Error(),
			})
			return
		}

		if err := validate.Struct(&request); err != nil {
			c.JSON(http.StatusBadRequest, ExchangeResponse{
				Msg: err.Error(),
			})
			return
		}

		amount, err := changeService.Exchange(exchangeRate.Currencies, request.Amount, request.Source, request.Target)

		if err != nil {
			c.JSON(http.StatusBadRequest, ExchangeResponse{
				Msg: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, ExchangeResponse{
			Msg:    "success",
			Amount: amount,
		})

	})

	route.Run(":80")
}
