package exchangerates

import (
	"encoding/json"
	"fmt"
	"strings"
)

// {
//   "base": "USD",
//   "date": "2022-04-14",
//   "rates": {
//     "EUR": 0.813399,
//     "GBP": 0.72007,
//     "JPY": 107.346001
//   },
//   "success": true,
//   "timestamp": 1519296206
// }

type Rate map[string]float64

type FixerAPIResponse struct {
	Base string `json:"base"`
	Date string `json:"date"`
	Rates Rate `json:"rate"`
	Success bool `json:"success"`
	Timestamp int64 `json:"timestamp"`
}

type FrankFurterAPIResponse struct {
	Amount float64 `json:"amount"`
	Base string `json:"base"`
	Date string `json:"date"`
	Rates Rate `json:"rates"`
}


func FormatAPIResponse(resp, url string) (Rate, error) {
	if strings.Contains(url, "frankfurter") {
		return formatFrankFurterResponse(resp)
	}
	return formatFixerAPIResponse(resp)
}

func formatFixerAPIResponse(resp string) (Rate, error) {
	var fixerResponse FixerAPIResponse
	err := json.Unmarshal([]byte(resp), &fixerResponse)
	if err != nil {
		fmt.Println("Error: converting api response to struct", err.Error())
		return nil, err
	}
	return fixerResponse.Rates, nil;
}

func formatFrankFurterResponse(resp string) (Rate, error) {
	var frankFurterResponse FrankFurterAPIResponse
	err := json.Unmarshal([]byte(resp), &frankFurterResponse)
	if err != nil {
		fmt.Println("Error: converting api response to struct", err.Error())
		return nil, err
	}
	return frankFurterResponse.Rates, nil;
}