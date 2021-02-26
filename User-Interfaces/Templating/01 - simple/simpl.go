package main

import (
	"net/http"
	"text/template"
	"time"
)

var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

var ecoTemplate = parseTemplateFile("template_2.html", "templates", funcMap) // for resources economy we put template parsing out and execute this once

func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}

type Page struct {
	Title, Content, Id string
	Date               time.Time
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Crab me once",
		Content: "Have fun with crabs and may be birds...",
		Id:      "Mr. Crab",
		Date:    time.Now(),
	}

	//t := parseTemplateFile("template_2.html", "templates", funcMap)
	ecoTemplate.Execute(w, p)
}

func parseTemplateFile(name, folder string, funcs template.FuncMap) *template.Template {
	t := template.Must(template.New(name).Funcs(funcs).ParseFiles(folder + "/" + name)) // New("...") must have name = the template file to work properly
	return t
}

func parseTemplateString(name, tpl string, funcs template.FuncMap) *template.Template {
	t := template.New(name)
	t.Funcs(funcs)
	t = template.Must(t.Parse(tpl))
	return t
}
