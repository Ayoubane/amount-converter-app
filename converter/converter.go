package converter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ConverterService struct {
}

type Currencies struct {
	Rates map[string]float64 `json:"rates"`
}

func (c ConverterService) GetRates(url string) (Currencies, error) {
	result := Currencies{}
	resp, err := http.Get(url)
	if err != nil {
		return result, err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &result); err != nil {
		return result, err
	}
	return result, nil
}

func (c ConverterService) Convert(currencies Currencies, value float64) map[string]float64 {
	result := map[string]float64{}
	for key, rate := range currencies.Rates {
		result[key] = rate * value
	}
	fmt.Printf("%v", result)
	return result
}
