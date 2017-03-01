//Example of a HTTP server with Go from: https://golang.org/doc/articles/wiki/#tmp_13
//A.Laszlo Lazuer

package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":8080", nil)
}
