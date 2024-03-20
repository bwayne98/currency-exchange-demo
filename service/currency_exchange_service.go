package service

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type CurrencyExchangeService struct{}

func (service *CurrencyExchangeService) Exchange(exchangeMap ExchangeMap, amountString string, source string, target string) (string, error) {

	amount, err := strToFloat(amountString)

	if err != nil {
		return "", err
	}

	if amount <= 0.0 {
		return "", errors.New("兌換金額小於或等於0")
	}

	inner, ok := exchangeMap[source]
	if !ok {
		return "", errors.New("來源貨幣種類 不支援")
	}

	power, ok := inner[target]
	if !ok {
		return "", errors.New("目標貨幣種類 不支援")
	}

	ret := math.Round(amount*power*100) / 100

	return fmt.Sprintf("%.2f", ret), nil

}

func strToFloat(amount string) (float64, error) {

	amount = strings.ReplaceAll(amount, ",", "")

	f, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		return 0.0, err
	}

	return f, nil
}
