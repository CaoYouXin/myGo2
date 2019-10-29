package main

import (
	"go-starter/ch1/lissajous"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL.Path, r.Proto)
	// for k, v := range r.Header {
	// 	fmt.Fprintf(w, "Header[%s]=%s\n", k, v)
	// }
	// fmt.Fprintf(w, "Host=%s\n", r.Host)
	// fmt.Fprintf(w, "RemoteAddr=%s\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
		return
	}
	for k, v := range r.Form {
		// fmt.Fprintf(w, "Form[%s]=%s\n", k, v)

		if k == "cycles" {
			explictCycles, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Print(err)
				return
			}

			lissajous.Lissajous(w, explictCycles)
		}
	}
}
