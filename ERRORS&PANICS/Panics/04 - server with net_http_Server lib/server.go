package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ashshaosh/scool/safemypanic"
)

func main() {
	safemypanic.Go(fearme)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	fearme() //panic(errors.New("just panic"))
}

func fearme() {
	fmt.Println("skdjfuhbwieufbwiu")
	panic(errors.New("fear my furr"))
}
