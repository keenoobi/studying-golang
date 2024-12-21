package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type ViewData struct {
	Hour int
}

func main() {

	data := ViewData{
		Hour: time.Now().Hour(),
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
