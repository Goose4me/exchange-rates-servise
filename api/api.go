package api

import (
	"fmt"

	"net/http"

	"gses2.app/rate"
)

const (
	genericErrorMsg            = "Something went wrong"
	currencyApiGenericErrorMsg = "Something went wrong with currency API"
)

func handleRate(w http.ResponseWriter, r *http.Request) {
	currencyRate, err := rate.GetCurrencyRateFor("usd", "uah")

	if err != nil {
		http.Error(w, currencyApiGenericErrorMsg, http.StatusBadRequest)

		return
	}

	if _, err := w.Write([]byte(currencyRate)); err != nil {
		http.Error(w, genericErrorMsg, http.StatusBadRequest)
	}
}

// StartServer starts the webserver to deliver jokes at /
func StartServer(port string) error {
	http.HandleFunc("/rate", handleRate)

	// Start the HTTP server on port 8080
	fmt.Printf("Server starting on port 8080...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return fmt.Errorf("error starting server: %s", err)
	}

	return nil
}
