package main

import (
	"net/http"
)

func main() {
	box := rice.MustFindBox("./files/")
	httpbox := box.HTTPBox()
	http.ListenAndServe(":8080", http.FileServer(httpbox))
}

// go get github.com/GeertJohan/rice
// rice embed-go // to embed files with virtual file-system
// go build
// and rice can help to prioretize some files for access
