package main

import (
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("./t/index.html", "./t/header.html"))
}

//Page ...
type Page struct {
	Title, Content string
}

func displayPage(rw http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Extermo!",
		Content: "Make no noise!",
	}
	t.ExecuteTemplate(rw, "index.html", p)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
