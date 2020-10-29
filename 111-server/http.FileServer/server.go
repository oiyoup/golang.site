package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/Users/oiyoup/Workspace/Documents")))
	http.ListenAndServe(":5000", nil)
}
