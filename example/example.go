package main

import (
	"log"
	"net/http"

	"github.com/geekfarmer/go-ipratelimit"
)

var limiter = ipratelimit.New(1, 5)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	if err := http.ListenAndServe(":8888", limiter.IPRateLimitMiddleware(mux)); err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	// Some very expensive database call
	w.Write([]byte("alles gut"))
}
