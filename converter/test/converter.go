package test

import "github.com/ayoubane/amount-converter-app/converter"

type MockConverterService struct {
}

func (c MockConverterService) GetRates(url string) (converter.Currencies, error) {
	result := converter.Currencies{}

	return result, nil
}

func (c MockConverterService) Convert(currencies converter.Currencies, value float64) map[string]float64 {
	result := map[string]float64{}
	return result
}
