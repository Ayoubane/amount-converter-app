package main

import (
	"fmt"
	"net/http"

	"github.com/ayoubane/amount-converter-app/api"
	"github.com/ayoubane/amount-converter-app/converter"
)

func main() {

	handler := api.ConverterHandler{ConverterService: converter.ConverterService{}}

	http.HandleFunc("/convert", handler.Convert)
	fmt.Println("Application started on port 8080")
	http.ListenAndServe(":8080", nil)

}
