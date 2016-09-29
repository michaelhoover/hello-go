package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("assets")))
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := http.ListenAndServe(addr, Log(http.DefaultServeMux)); err != nil {
		log.Fatal(err)
	}
}
