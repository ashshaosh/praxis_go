package main

import (
	"fmt"
	"net/http"

	fs "github.com/Masterminds/go-fileserver"
)

func main() {
	fs.NotFoundHandler = func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Contemt-Type", "text/plain; charset=utf-8")
		fmt.Println(w, "Requested page could not be found")
	}
	dir := http.Dir("./files")
	http.ListenAndServe(":8080", fs.FileServer(dir))
}
