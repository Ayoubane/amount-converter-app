package main

import (
	"fmt"
	"net/http"

	"github.com/ayoubane/amount-converter-app/api"
)

func main() {

	http.HandleFunc("/convert", api.Handler)
	fmt.Println("Application started on port 8080")
	http.ListenAndServe(":8080", nil)

}
