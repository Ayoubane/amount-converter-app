package converter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	c := ConverterService{}
	cur := Currencies{Rates: map[string]float64{"USD": 1.100, "MYR": 23.94, "JYP": 253.00}}
	result := c.Convert(cur, 15.0)
	test := map[string]float64{"USD": 16.5, "MYR": 359.1, "JYP": 3795.0}
	if ok := reflect.DeepEqual(result, test); !ok {
		t.Error("Result Unexpected")
	}
}

func TestGetRates(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"base":"EUR",
				"date":"2016-10-12",
				"rates":
					{"AUD":1.4553,
					"BGN":1.9558,
					"BRL":3.5256,
					"CAD":1.4598}
				}`)
	}))
	defer server.Close()
	c := ConverterService{}
	rates := map[string]float64{"AUD": 1.4553,
		"BGN": 1.9558,
		"BRL": 3.5256,
		"CAD": 1.4598}
	result, _ := c.GetRates(server.URL)
	if ok := reflect.DeepEqual(rates, result.Rates); !ok {
		t.Error("Unexpected Error")
	}
}
