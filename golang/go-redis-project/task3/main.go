package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// ping, err := client.Ping(context.Background()).Result()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(ping)

	http.HandleFunc("/cache", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var data Data
		if err = json.Unmarshal(body, &data); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		err = client.Set(context.Background(), data.Key, data.Value, 0).Err()
		if err != nil {
			http.Error(w, "Failed to save data to Redis", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"status": "success", "message": "Data saved to Redis"}

		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
