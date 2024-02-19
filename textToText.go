package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func textToText(payload Payload, api string, token string) ([]interface{}, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", api, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return convertToInterfaceSlice(result), err
}

func convertToInterfaceSlice(data interface{}) []interface{} {
	switch v := data.(type) {
	case []interface{}:
		slice := make([]interface{}, len(v))
		for i, elem := range v {
			slice[i] = convertToInterfaceSlice(elem)
		}
		return slice
	case map[string]interface{}:
		return []interface{}{v}
	default:
		return []interface{}{data}
	}
}
