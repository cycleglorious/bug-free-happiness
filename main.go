package main

import (
	"fmt"
	"net/http"
)

func Add(a string, b string) string {
	return fmt.Sprintf("%s%sa", a, b)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	w.Write([]byte(string(Add(a, b))))
}

func main() {
	http.HandleFunc("GET /add", AddHandler)
	http.ListenAndServe(":8080", nil)
}
