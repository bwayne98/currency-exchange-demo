package main

import (
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

	service := service.CurrencyExchangeService{}
	exchangeMap := map[string]map[string]float64{
		"TWD": {
			"TWD": 1,
			"JPY": 3.669,
			"USD": 0.03281,
		},
		"JPY": {
			"TWD": 0.26956,
			"JPY": 1,
			"USD": 0.00885,
		},
		"USD": {
			"TWD": 30.444,
			"JPY": 111.801,
			"USD": 1,
		},
	}

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

		amount, err := service.Exchange(exchangeMap, request.Amount, request.Source, request.Target)

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
