package main

import (
	"fmt"
	"log"
    "net/http"
)

func main() {
    http.Handle("/", new(testHandler))

    err := http.ListenAndServe(":5000", nil)
    if err != nil {
        log.Fatal(err)
    }
}

type testHandler struct {
    http.Handler
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    str := "Your Request Path is " + req.URL.Path
    write, err := w.Write([]byte(str))
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(write)
}
