package main

import (
	"net/http"
)

// sort
// surface & eval
// decode xml

type StringWriter interface {
	WriteString(string) (int, error)
}

func main() {
	http.HandleFunc("/", func(wr http.ResponseWriter, r *http.Request) {
		if wr, ok := wr.(StringWriter); ok {
			wr.WriteString("世界你好👋")
		}
	})
	http.ListenAndServe("localhost:8000", nil)
}
