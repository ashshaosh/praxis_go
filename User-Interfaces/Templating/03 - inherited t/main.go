package main

import (
	"net/http"
	"text/template"
)

var t map[string]*template.Template

func init() {
	t = make(map[string]*template.Template)
	temp := template.Must(template.ParseFiles("./t/base.html", "./t/page.html"))
	t["page.html"] = temp
	temp = template.Must(template.ParseFiles("./t/base.html", "./t/user.html"))
	t["user.html"] = temp
}

type Page struct {
	Title, Content string
}

type User struct {
	Username, Name string
}

func displayPage(rw http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "HI TITLER",
		Content: "Stardust marionette wach the watch",
	}
	t["page.html"].ExecuteTemplate(rw, "base", p)
}

func displayUser(rw http.ResponseWriter, r *http.Request) {
	u := &User{
		Username: "lolloop",
		Name:     "Andy Titler",
	}
	t["user.html"].ExecuteTemplate(rw, "base", u)
}

func main() {
	http.HandleFunc("/user", displayUser)
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
