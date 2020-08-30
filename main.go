package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

//Page struct for displaying page
type Page struct {
	Title, Content string
}

var t map[string]*template.Template

func init() {
	t = make(map[string]*template.Template)
	temp := template.Must(template.ParseFiles("./templates/index.html", "./templates/head.html", "./templates/hello.html"))
	t["hello.html"] = temp
	temp = template.Must(template.ParseFiles("./templates/index.html", "./templates/head.html", "./templates/languages.html"))
	t["lang.html"] = temp
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Speech to lang",
		Content: "Chatbot that helps you with learning",
	}
	var b bytes.Buffer
	err := t["hello.html"].ExecuteTemplate(&b, "index", p)
	if err != nil {
		fmt.Fprint(w, "A error occured.")
		return
	}
	b.WriteTo(w)
}

func createEngRoom(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "English room",
		Content: "Hello! Lets speak",
	}
	var b bytes.Buffer
	err := t["lang.html"].ExecuteTemplate(&b, "index", p)
	if err != nil {
		fmt.Fprint(w, "A error occured.")
		return
	}
	b.WriteTo(w)
	fmt.Println("user entered english room")
}

func createRusRoom(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Русский зал",
		Content: "Привет! Давай начнем говорить",
	}
	var b bytes.Buffer
	err := t["lang.html"].ExecuteTemplate(&b, "index", p)
	if err != nil {
		fmt.Fprint(w, "A error occured.")
		return
	}
	b.WriteTo(w)
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
