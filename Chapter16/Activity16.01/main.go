package main

import (
	"fmt"
	"net/http"
    "log"
)

type PageWithCounter struct {
    counter int
    content string
    heading string
}

func (p *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    p.counter++
    msg := fmt.Sprintf("Header: %s, Content: %s, page views: %d", p.content, p.heading, p.counter)
    w.Write([]byte(msg))
}

func main() {
    hello := PageWithCounter{heading: "hello world", content: "This is the main page"}
    chapter1 := PageWithCounter{heading: "Chapter 1", content: "This is the first chapter"}
    chapter2 := PageWithCounter{heading: "Chapter 2", content: "This is the second chapter"}
    
    http.Handle("/", &hello)
    http.Handle("/chapter1", &chapter1)
    http.Handle("/chapter2", &chapter2)
    log.Fatal(http.ListenAndServe(":8080", nil))
}