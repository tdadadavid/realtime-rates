package utils

import (
	"io"
	"net/http"
)

type RequestParams struct {
	Url string
	Key  string
}

func HandleRequest(opt RequestParams) (string, error) {

	req, err := http.NewRequest("GET", opt.Url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("apikey", opt.Key)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}