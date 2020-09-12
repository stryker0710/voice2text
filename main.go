package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("soundBlob")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func main() {
	mux := http.NewServeMux()

	/*muxRoutes:=make(map[string]http.Handler)
	muxRoutes["images"]=http.FileServer(http.Dir("./images"))
	muxRoutes["js"]=http.FileServer(http.Dir("./js"))
	muxRoutes["css"]=http.FileServer(http.Dir("./css"))*/
	route := "public"
	server := http.FileServer(http.Dir("./public"))

	mux.Handle("/"+route+"/", http.StripPrefix("/"+route, server))

	mux.HandleFunc("/", displayPage)
	mux.HandleFunc("/english/", createEngRoom)
	mux.HandleFunc("/russian/", createRusRoom)
	mux.HandleFunc("/upload/", uploadHandler)
	err := http.ListenAndServe(":9090", mux) // set listen port
	log.Fatal(err)

}
