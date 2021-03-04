package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type cacheFile struct { // struct for file
	content io.ReadSeeker
	modTime time.Time
}

var cache map[string]*cacheFile // slice of files in cache
var mutex = new(sync.RWMutex)   // lock cache for parallel changes

func main() {
	cache = make(map[string]*cacheFile)
	http.HandleFunc("/", serveFiles)
	http.ListenAndServe(":8080", nil)
}

func serveFiles(res http.ResponseWriter, req *http.Request) {
	mutex.RLock()
	v, found := cache[req.URL.Path] // get if file already cached
	mutex.RUnlock()
	if !found { // file isn't cached
		mutex.Lock()
		defer mutex.Unlock()

		fileName := "./files" + req.URL.Path
		f, err := os.Open(fileName) // open file for caching
		defer f.Close()             // set delay for closing file
		if err != nil {             // opening errors
			http.NotFound(res, req)
			return
		}
		var b bytes.Buffer      // make buffer
		_, err = io.Copy(&b, f) // copy file into buffer
		if err != nil {         // copyng errors
			http.NotFound(res, req)
			return
		}
		r := bytes.NewReader(b.Bytes()) // put bytes into Reader

		info, _ := f.Stat()
		v := &cacheFile{ // fill object for caching
			content: r,
			modTime: info.ModTime(),
		}
		cache[req.URL.Path] = v // add object to slice
	}
	http.ServeContent(res, req, req.URL.Path, v.modTime, v.content) // get file from cache
}
