package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	data []string
	mu   sync.Mutex
}

func (s *Server) handlePostData(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Data string `json:"data"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, request.Data)
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) handleGetData(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	response := struct {
		Data []string `json:"data"`
	}{
		Data: s.data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	server := &Server{data: make([]string, 0)}

	// Создаем маршрутизатор с использованием chi
	router := chi.NewRouter()

	// Добавляем middleware для логирования и обработки ошибок
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Определяем маршруты
	router.Post("/data", server.handlePostData)
	router.Get("/data", server.handleGetData)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}
