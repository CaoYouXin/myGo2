package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var counter int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", count)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func count(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	fmt.Fprintf(w, "Server Hit Count: %d\n", counter)
	mu.Unlock()
}
