package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

//Page struct for displaying page
type Page struct {
	Title, Content string
}

//LangPage struct for displaying page
type LangPage struct {
	Room, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Speech to lang",
		Content: "Chatbot that helps you with learning",
	}
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, p)
}

func createEngRoom(w http.ResponseWriter, r *http.Request) {
	p := &LangPage{
		Room:    "English room",
		Content: "Hello! Lets speak",
	}
	t := template.Must(template.ParseFiles("templates/languages.html"))
	t.Execute(w, p)
	fmt.Println("user entered english room")
}

func createRusRoom(w http.ResponseWriter, r *http.Request) {
	p := &LangPage{
		Room:    "Русский зал",
		Content: "Привет! Давай начнем говорить",
	}
	t := template.Must(template.ParseFiles("templates/languages.html"))
	t.Execute(w, p)
	fmt.Println("user entered russian room")
}

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./images"))
	mux.Handle("/images/", http.StripPrefix("/images", fileServer))
	jsServer := http.FileServer(http.Dir("./js"))
	mux.Handle("/js/", http.StripPrefix("/js", jsServer))
	cssServer := http.FileServer(http.Dir("./css"))
	mux.Handle("/css/", http.StripPrefix("/css", cssServer))
	mux.HandleFunc("/", displayPage)
	mux.HandleFunc("/english/", createEngRoom)
	mux.HandleFunc("/russian/", createRusRoom)
	err := http.ListenAndServe(":9090", mux) // set listen port
	log.Fatal(err)

}
