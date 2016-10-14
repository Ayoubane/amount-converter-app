package main

import (
	"net/http"

	"github.com/ayoubane/amount-converter-app/api"
	"github.com/ayoubane/amount-converter-app/converter"
)

func init() {

	handler := api.ConverterHandler{ConverterService: converter.ConverterService{}}

	http.HandleFunc("/convert", handler.Convert)

}
