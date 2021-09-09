package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var welcome = template.Must(template.ParseFiles("templates/welcome.html"))
var index = template.Must(template.ParseFiles("templates/index.html"))
var gallery = template.Must(template.ParseFiles("templates/gallery.html"))

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	welcome.Execute(w, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, nil)
}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
	gallery.Execute(w, nil)
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

	mux.HandleFunc("/", welcomeHandler)
	mux.HandleFunc("/welcome", indexHandler)
	mux.HandleFunc("/gallery", galleryHandler)
	mux.Handle("/public/", http.StripPrefix("/public", fs))
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
