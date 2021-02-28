package main

import "net/http"

func main() {
	dir := http.Dir("./files")
	prefiks := http.StripPrefix("/static/", http.FileServer(dir))
	http.Handle("/static/", prefiks)

	//http.HandleFunc("/", homepage)
	http.ListenAndServe(":8080", http.FileServer(dir))
}
