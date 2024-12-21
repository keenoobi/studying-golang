package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "http://server:8080/data" // Адрес сервера

	// Отправляем POST-запрос для сохранения данных
	data := []string{"Hello", "World", "from", "client"}
	for _, item := range data {
		payload := map[string]string{"data": item}
		jsonPayload, _ := json.Marshal(payload)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
		if err != nil {
			fmt.Printf("Error sending data: %v\n", err)
			return
		}
		resp.Body.Close()
	}

	// Отправляем GET-запрос для получения данных
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting data: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var response struct {
		Data []string `json:"data"`
	}
	json.NewDecoder(resp.Body).Decode(&response)

	fmt.Printf("Received data from server: %v\n", response.Data)
}
