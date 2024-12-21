package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Message struct {
	Topic   string `json:"topic"`
	Content string `json:"Content"`
}

func main() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only Post requests are allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read reques body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var msg Message
		if err = json.Unmarshal(body, &msg); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &msg.Topic, Partition: kafka.PartitionAny},
			Value:          []byte(msg.Content),
		}, nil)
		if err != nil {
			http.Error(w, "Failed to send message to Kafka", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"status": "success", "message": "Message sent to Kafka"}

		json.NewEncoder(w).Encode(response)

	})

	fmt.Println("Server is listening on http://localhost:8181")
	http.ListenAndServe(":8181", nil)
}
