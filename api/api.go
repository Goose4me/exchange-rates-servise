package api

import (
	"fmt"

	"net/http"

	"gses2.app/database"
	"gses2.app/models"
	"gses2.app/rate"
)

// Generic messages for user
const (
	genericErrorMsg     = "Something went wrong"
	currencyApiErrorMsg = "Something went wrong with currency API"
)

func handleRate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		currencyRate, err := rate.GetCurrencyRateFor("usd", "uah")

		if err != nil {
			http.Error(w, currencyApiErrorMsg, http.StatusBadRequest)

			return
		}

		if _, err := w.Write([]byte(currencyRate)); err != nil {
			http.Error(w, genericErrorMsg, http.StatusInternalServerError)
		}
	}
}

func handleSubscription(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		// Parse form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, genericErrorMsg, http.StatusInternalServerError)
			return
		}

		email := r.FormValue("email")
		if email == "" {
			http.Error(w, "\"email\" parameter is empty", http.StatusBadRequest)

			return
		}

		var subscriber *models.Subscriber
		result := database.DB.Db.Model(models.Subscriber{}).Limit(1).Find(&subscriber, "Email = ?", email)

		if result.Error != nil {
			http.Error(w, "", http.StatusInternalServerError)

			return
		}

		// Check if exists
		if result.RowsAffected > 0 {
			http.Error(w, "", http.StatusConflict)

			return
		}

		subscriber = new(models.Subscriber)
		subscriber.Email = email

		database.DB.Db.Create(&subscriber)

		w.WriteHeader(http.StatusOK)
	}
}

// StartServer starts the webserver to deliver jokes at /
func StartServer(port string) error {
	http.HandleFunc("/api/rate", handleRate)
	http.HandleFunc("/api/subscribe", handleSubscription)

	// Start the HTTP server on port
	fmt.Printf("Server starting on port " + port + "...\n")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return fmt.Errorf("error starting server: %s", err)
	}

	return nil
}
