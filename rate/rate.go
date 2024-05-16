package rate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const (
	primaryApiBaseUrl   = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/"
	secondaryApiBaseUrl = "https://latest.currency-api.pages.dev/v1/currencies/"
)

func getRate(responseData []byte, from string, to string) (string, error) {
	// API response should have format {"date": "${yyyy-mm-dd}", "${currency}": {"${currency}": ${rate}, ...}}
	var result map[string]interface{}

	err := json.Unmarshal(responseData, &result)
	if err != nil {
		fmt.Printf("Error parsing JSON: %s", err)
		return "", fmt.Errorf("error parsing JSON: %s", err)
	}

	currencies := result[from].(map[string]interface{})
	return strconv.FormatFloat(currencies[to].(float64), 'f', 3, 64), nil
}

func callSecondaryUrl(from *string) (*http.Response, error) {
	response, err := http.Get(secondaryApiBaseUrl + *from + ".min.json")
	fmt.Printf("[SecondaryURL] The HTTP request failed with fallback url error %s\n", err)
	if err != nil {
		return response, fmt.Errorf("the HTTP request failed with error %s", err)
	}

	// Handle bad responses from secondary URL. Most posssibly "404 Not Found"
	if response.StatusCode != 200 {
		fmt.Printf("[SecondaryURL] Currency API call from secondary URL was not succeed. Status code: %d\n", response.StatusCode)
		return response, fmt.Errorf("currency API call was not succeed. Status code: %d", response.StatusCode)
	}

	return response, err
}

func GetCurrencyRateFor(from string, to string) (string, error) {
	response, err := http.Get(primaryApiBaseUrl + from + ".min.json")
	if err != nil {
		fmt.Printf("[PrimaryURL] The HTTP request failed with primary url error %s\n", err)

		// Use the secondary URL if the primary haven't worked
		response, err = callSecondaryUrl(&from)

		if err != nil {
			return "", err
		}
	}

	defer response.Body.Close()

	// Handle bad responses from primary URL. Most posssibly "404 Not Found"
	if response.StatusCode != 200 {
		fmt.Printf("[PrimaryURL] Currency API call was not succeed. Status code: %d", response.StatusCode)

		// Use the secondary URL if the primary haven't worked
		response, err = callSecondaryUrl(&from)

		if err != nil {
			return "", err
		}
	}

	// Read the response body
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Reading response body failed with error %s\n", err)
		return "", fmt.Errorf("reading response body failed with error %s", err)
	}

	return getRate(responseData, from, to)
}
