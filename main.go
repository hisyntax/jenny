package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fs := http.FileServer(http.Dir("public"))
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.Handle("/public/", http.StripPrefix("/public", fs))
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
