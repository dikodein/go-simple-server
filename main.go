package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", middleware(index()))

	http.ListenAndServe(":8080", mux)
}

func index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I'm Awake\n"))
	}
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You'l see me first\n"))
		next.ServeHTTP(w, r)
	}
}
