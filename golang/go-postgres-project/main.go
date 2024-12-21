package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable"

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to database")
}

func createUser(w http.ResponseWriter, req *http.Request) {
	var user User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	db.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var user User
	db.First(&user, id)
	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Index page")
}

func main() {
	initDB()
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
