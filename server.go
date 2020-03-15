package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.RequestURI == "/" {
			req.RequestURI = "/index.html"
		}
		file, err := os.Open(filepath.Join("./site", req.RequestURI))
		if err == nil {
			io.Copy(w, file)
		}
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}
