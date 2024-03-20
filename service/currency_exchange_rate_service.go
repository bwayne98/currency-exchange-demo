package service

import (
	"encoding/json"
)

type CurencyExchngeRateService struct{}

type ExchangeMap = map[string]map[string]float64

type ExchageRate struct {
	Currencies ExchangeMap `json:"currencies"`
}

func (service CurencyExchngeRateService) Fetch() (rate *ExchageRate, err error) {
	raw := getRawData()

	if err := json.Unmarshal([]byte(raw), &rate); err != nil {
		return nil, err
	}

	return
}

func getRawData() string {
	return "{\n\"currencies\": {\n\"TWD\": {\n\"TWD\": 1,\n\"JPY\": 3.669,\n\"USD\": 0.03281\n},\n\"JPY\": {\n\"TWD\": 0.26956,\n\"JPY\": 1,\n\"USD\": 0.00885\n},\n\"USD\": {\n\"TWD\": 30.444,\n\"JPY\": 111.801,\n\"USD\": 1\n}\n}\n}"
}
