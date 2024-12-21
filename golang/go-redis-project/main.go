package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Person struct {
	ID         string
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ping)

	elliotID := uuid.NewString()
	jsonString, err := json.Marshal(Person{
		ID:         elliotID,
		Name:       "Elliot",
		Age:        30,
		Occupation: "Staff Software Engineer",
	})
	if err != nil {
		fmt.Printf("faile to marshal: %s", err.Error())
		return
	}

	elliotKey := fmt.Sprintf("person:%s", elliotID)

	err = client.Set(context.Background(), elliotKey, jsonString, 0).Err()
	if err != nil {
		fmt.Printf("Failed to set value in the redis instance: %s", err.Error())
		return
	}

	val, err := client.Get(context.Background(), elliotKey).Result()
	if err != nil {
		fmt.Printf("failed to get value from redis: %s", err.Error())
	}

	fmt.Printf("value retrieved from redis: %s\n", val)

	// ctx := context.Background()
	// pong, err := client.Ping(ctx).Result()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Redis connection:", pong)
}
