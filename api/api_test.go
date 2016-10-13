package api

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"fmt"

	"github.com/ayoubane/amount-converter-app/converter"
	"github.com/ayoubane/amount-converter-app/converter/test"
)

func TestConverterHandler_Convert(t *testing.T) {
	mockConverterService := test.MockConverterService{}
	mockCurrencies := converter.Currencies{}
	mockConverterService.On("GetRates", "https://api.fixer.io/latest").Return(mockCurrencies, nil)
	mockConverterService.On("Convert", mockCurrencies, 50.0).Return(map[string]float64{"EUR": 50.0, "USD": 55.0})

	converterHandler := ConverterHandler{mockConverterService}
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/converter?amount=50.0", nil)
	converterHandler.Convert(response, request)
	expected := `{
    "EUR": 50,
    "USD": 55
}`
	if response.Body.String() != expected {
		t.Errorf("Erreur sur la conversion. expected: %v got: %v", expected, response.Body.String())
	}

}

func TestConverterHandler_Convert_InvalidAmount(t *testing.T) {
	mockConverterService := test.MockConverterService{}
	mockConverterService.On("GetRates", "https://api.fixer.io/latest").Return(converter.Currencies{}, nil)

	converterHandler := ConverterHandler{mockConverterService}
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/converter?amount=50.0ERR", nil)
	converterHandler.Convert(response, request)
	if response.Code != http.StatusBadRequest {
		t.Error("Status bad request attendu")
	}
}

func TestConverterHandler_Convert_RatesUnreachable(t *testing.T) {
	mockConverterService := test.MockConverterService{}
	mockConverterService.On("GetRates", "https://api.fixer.io/latest").Return(converter.Currencies{}, fmt.Errorf("Url rates invalide"))
	converterHandler := ConverterHandler{mockConverterService}
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/converter?amount=50.0ERR", nil)
	converterHandler.Convert(response, request)
	if response.Code != http.StatusInternalServerError {
		t.Error("Status Internal Server attendu")
	}

}
