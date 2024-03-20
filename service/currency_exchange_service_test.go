package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExchange(t *testing.T) {
	testCase := []struct {
		name      string
		payload   func() (exchangeMap ExchangeMap, amountString string, source string, target string)
		checkFunc func(t *testing.T, result string, err error)
	}{
		{
			name: "source 不支援",
			payload: func() (exchangeMap ExchangeMap, amountString string, source string, target string) {
				exchangeMap = ExchangeMap{
					"TWD": {
						"USD": 0.1,
					},
				}

				amountString = "1.0"
				source = "USD"

				return
			},
			checkFunc: func(t *testing.T, result string, err error) {
				require.NotNil(t, err)
				require.Equal(t, err.Error(), "來源貨幣種類 不支援")
				require.Equal(t, result, "")
			},
		},
		{
			name: "target 不支援",
			payload: func() (exchangeMap ExchangeMap, amountString string, source string, target string) {
				exchangeMap = ExchangeMap{
					"TWD": {
						"USD": 0.1,
					},
				}

				amountString = "1.0"
				source = "TWD"
				target = "JPD"

				return
			},
			checkFunc: func(t *testing.T, result string, err error) {
				require.NotNil(t, err)
				require.Equal(t, err.Error(), "目標貨幣種類 不支援")
				require.Equal(t, result, "")
			},
		},
		{
			name: "amount 小於或等於0",
			payload: func() (exchangeMap ExchangeMap, amountString string, source string, target string) {
				exchangeMap = ExchangeMap{
					"TWD": {
						"USD": 0.1,
					},
				}

				amountString = "-1.0"
				source = "TWD"
				target = "USD"

				return
			},
			checkFunc: func(t *testing.T, result string, err error) {
				require.NotNil(t, err)
				require.Equal(t, err.Error(), "兌換金額小於或等於0")
				require.Equal(t, result, "")
			},
		},
		{
			name: "amount 無法轉換成float64",
			payload: func() (exchangeMap ExchangeMap, amountString string, source string, target string) {
				exchangeMap = ExchangeMap{
					"TWD": {
						"USD": 0.1,
					},
				}

				amountString = "abc"
				source = "TWD"
				target = "USD"

				return
			},
			checkFunc: func(t *testing.T, result string, err error) {
				require.NotNil(t, err)
				require.Equal(t, err.Error(), "金額格式錯誤")
				require.Equal(t, result, "")
			},
		},
		{
			name: "成功轉換並進位",
			payload: func() (exchangeMap ExchangeMap, amountString string, source string, target string) {
				exchangeMap = ExchangeMap{
					"TWD": {
						"USD": 0.115,
					},
				}

				amountString = "1"
				source = "TWD"
				target = "USD"

				return
			},
			checkFunc: func(t *testing.T, result string, err error) {
				require.Nil(t, err)
				require.Equal(t, result, "0.12")
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			service := CurrencyExchangeService{}

			exchangeMap, amountString, source, target := tc.payload()

			result, err := service.Exchange(exchangeMap, amountString, source, target)

			tc.checkFunc(t, result, err)
		})
	}
}
