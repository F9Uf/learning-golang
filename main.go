package main

import "net/http"

func main() {
	// :8080 => 0.0.0.0:8080
	http.ListenAndServe(":8080", http.HandlerFunc(router))
}

// http Handler
// vs code just write `hand`, it auto generate function
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	w.Write([]byte("<h1>Hello</>"))
}

func router(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		indexHandler(w, r)
	} else if r.URL.Path == "/about" {
		aboutHandler(w, r)
	} else {
		notFoundHandler(w, r)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about page"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not found"))
}
