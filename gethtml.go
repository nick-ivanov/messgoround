package main

import "net/http"

func getHTML(url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", err
	}

	b := make([]byte, 999999)
	res.Body.Read(b)

	return string(b), nil
}
