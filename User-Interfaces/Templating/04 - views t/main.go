package main

import (
	"bytes"
	"html/template"
	"net/http"
)

var t *template.Template
var qc template.HTML

func init() {
	t = template.Must(template.ParseFiles("./t/index.html", "./t/quote.html"))
}

type Page struct {
	Title   string
	Content template.HTML
}

type Quote struct {
	Quote, Person string
}

func main() {
	q := &Quote{
		Quote:  `We must not allow other people’s limited perceptions to define us.`,
		Person: "Virginia Satir",
	}
	var b bytes.Buffer
	t.ExecuteTemplate(&b, "quote.html", q)
	qc = template.HTML(b.String())

	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}

func displayPage(rw http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Quote",
		Content: qc,
	}
	t.ExecuteTemplate(rw, "index.html", p)
}
