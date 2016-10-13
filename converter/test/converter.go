package test

import (
	"github.com/ayoubane/amount-converter-app/converter"
	"github.com/stretchr/testify/mock"
)

type MockConverterService struct {
	mock.Mock
}

func (c MockConverterService) GetRates(url string) (converter.Currencies, error) {
	args := c.Called(url)
	return args.Get(0).(converter.Currencies), args.Error(1)

}

func (c MockConverterService) Convert(currencies converter.Currencies, value float64) map[string]float64 {
	args := c.Called(currencies, value)
	return args.Get(0).(map[string]float64)
}
