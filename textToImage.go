package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func textToImageQuery(url string, token string, payload Payload) ([]byte, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	imageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return imageBytes, nil
}
