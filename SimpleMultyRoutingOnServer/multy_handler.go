package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", hello)      // when we multyroute
	http.HandleFunc("/goodbye/", goodbye) // at first list certain routes
	http.HandleFunc("/", homePage)        // then default one last to handle 404
	http.ListenAndServe(":8080", nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Brave Bear's Bare Bones Brothers"
	}
	fmt.Fprint(res, "Hello, i'm ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Babaykinson"
	}
	fmt.Fprint(res, "By, ", name)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "HOME")
}
