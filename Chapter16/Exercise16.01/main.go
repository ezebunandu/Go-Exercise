package main

import (
	"log"
	"net/http"
)

type hello struct {}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    mgs := "<h1>Hello World</h1>"
    w.Write([]byte(mgs))
}

func main() {
    log.Fatal(http.ListenAndServe(":8080", hello{}))
}