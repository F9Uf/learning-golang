package main

import (
	"fmt"
	"net/http"
)

func main() {
	// :8080 => 0.0.0.0:8080
	m := http.NewServeMux()
	m.Handle("/", http.HandlerFunc(indexHandler))
	m.Handle("/about", http.HandlerFunc(aboutHandler))

	http.ListenAndServe(":8080", chainMiddlewares(
		m1,
		m2,
	)(m))

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

// middleware
type middleware func(http.Handler) http.Handler

func chainMiddlewares(ms ...middleware) middleware {
	return func(h http.Handler) http.Handler {
		for i := len(ms) - 1; i >= 0; i-- {
			h = ms[i](h)
		}
		return h
	}
}

func m1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before m1")
		h.ServeHTTP(w, r)
		fmt.Println("after m1")
	})
}

func m2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before m2")
		h.ServeHTTP(w, r)
		fmt.Println("after m2")
	})
}

// private=lowercase first char name, public=uppercase first char name
