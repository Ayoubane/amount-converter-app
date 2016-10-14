package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"google.golang.org/appengine"

	"log"

	"github.com/ayoubane/amount-converter-app/converter"
	"golang.org/x/net/context"
)

type ConverterHandler struct {
	ConverterService ConverterServiceInterface
}

type ConverterServiceInterface interface {
	GetRates(ctx context.Context, url string) (converter.Currencies, error)
	Convert(currencies converter.Currencies, value float64) map[string]float64
}

func (h ConverterHandler) Convert(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	result, err := h.ConverterService.GetRates(ctx, "https://api.fixer.io/latest")
	if err != nil {
		log.Printf("Error : %v", err)
		http.Error(w, "Probleme sur la lecture du JSon rates", http.StatusInternalServerError)
		return
	}
	/*for key, rate := range result.Rates {
		fmt.Printf("%s _ %f | ", key, rate)
	}*/
	amount, err := strconv.ParseFloat(r.URL.Query().Get("amount"), 64)
	if err != nil {
		http.Error(w, "Probleme dans le parsing", http.StatusBadRequest)
		return
	}
	rs := h.ConverterService.Convert(result, amount)
	//fmt.Printf("%v", rs)
	b, _ := json.Marshal(rs)
	w.Header().Add("Content-Type", "application/json")
	w.Write(indentJSON(b))
	//fmt.Fprintf(w, "\n %v", rs)
}

func indentJSON(src []byte) []byte {
	dest := new(bytes.Buffer)
	err := json.Indent(dest, src, "", "    ")
	if err != nil {
		return nil
	}
	return dest.Bytes()
}
