package api

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"

	"github.com/ayoubane/amount-converter-app/converter"
	"github.com/ayoubane/amount-converter-app/converter/test"
)

var _ = Describe("Convert", func() {
	var (
		mockConverterService test.MockConverterService
	)

	BeforeEach(func() {
		mockConverterService = test.MockConverterService{}
	})

	Describe("Converting amount to currencies", func() {
		Context("With correct amount", func() {
			It("should return valid expected JSON", func() {
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
				Expect(response.Body.String()).To(Equal(expected))
			})
		})

		Context("With incorrect amount", func() {
			It("should return BadRequest error", func() {
				mockConverterService.On("GetRates", "https://api.fixer.io/latest").Return(converter.Currencies{}, nil)
				converterHandler := ConverterHandler{mockConverterService}
				response := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", "/converter?amount=50.0ERR", nil)
				converterHandler.Convert(response, request)
				Expect(response.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Context("With incorrect URL", func() {
			It("should return InternalServerError", func() {
				mockConverterService.On("GetRates", "https://api.fixer.io/latest").Return(converter.Currencies{}, fmt.Errorf("Url rates invalide"))
				converterHandler := ConverterHandler{mockConverterService}
				response := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", "/converter?amount=50.0ERR", nil)
				converterHandler.Convert(response, request)
				Expect(response.Code).To(Equal(http.StatusInternalServerError))
			})
		})
	})
})
