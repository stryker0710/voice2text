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

	muxRoutes:=make(map[string]http.Handler)
	muxRoutes["images"]=http.FileServer(http.Dir("./images"))
	muxRoutes["js"]=http.FileServer(http.Dir("./js"))
	muxRoutes["css"]=http.FileServer(http.Dir("./css"))

	for route,server:=range muxRoutes{
		mux.Handle("/"+route+"/",http.StripPrefix("/"+route, server))
	}

	mux.HandleFunc("/", displayPage)
	mux.HandleFunc("/english/", createEngRoom)
	mux.HandleFunc("/russian/", createRusRoom)
	err := http.ListenAndServe(":9090", mux) // set listen port
	log.Fatal(err)

}
