package main

import (
	"bytes"
	"html/template"
	"net/http"
)

var t *template.Template
var qc, ac template.HTML

func init() { // get template to know html files needed
	t = template.Must(template.ParseFiles("./t/index.html", "./t/quote.html", "./t/article.html"))
}

//Page struct to describe page template
type Page struct {
	Title        string
	PageContent  template.HTML
	QuoteContent template.HTML
}

//Article struct for article inside page
type Article struct {
	Title, Autor, PublicationDate, ArticleContent string
}

// Quote struct for quote block inside page
type Quote struct {
	Quote, Person string
}

func main() {
	q := &Quote{ // var for hold quote parameters
		Quote:  `We must not allow other people’s limited perceptions to define us.`,
		Person: "Virginia Satir",
	}
	artcl := &Article{ // var for hold article parameters
		Title:           "Pagoda (data structure)",
		Autor:           "Shri Shrimad Bhaktevedanta Svami Brapkhupada",
		PublicationDate: "22.02.20",
		ArticleContent:  `In computer science, a pagoda is a priority queue implemented with a variant of a binary tree. The root points to its children, as in a binary tree. Every other node points back to its parent and down to its leftmost (if it is a right child) or rightmost (if it is a left child) descendant leaf. The basic operation is merge or meld, which maintains the heap property. An element is inserted by merging it as a singleton. The root is removed by merging its right and left children. Merging is bottom-up, merging the leftmost edge of one with the rightmost edge of the other.`,
	}
	var b bytes.Buffer

	t.ExecuteTemplate(&b, "article.html", artcl) // parse var with template for article into buffer
	ac = template.HTML(b.String())               // fill var with string from buffer
	b.Reset()                                    // clear buffer
	t.ExecuteTemplate(&b, "quote.html", q)       // parse into buffer
	qc = template.HTML(b.String())               // fill var ...
	b.Reset()

	http.HandleFunc("/", displayPage) // assign function in case of connection
	http.ListenAndServe(":8080", nil) // start listening for connection
}

func displayPage(rw http.ResponseWriter, r *http.Request) {
	p := &Page{ // make var to hold whole page content
		Title:        "Test Quote", // assign page itle
		PageContent:  ac,           // post var to parameter
		QuoteContent: qc,
	}
	t.ExecuteTemplate(rw, "index.html", p) // parse var with index.html template
}
