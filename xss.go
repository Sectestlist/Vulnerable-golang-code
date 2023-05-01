package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<html><body>Hello, " + r.URL.Query().Get("name") + "</body></html>"))
	})
	http.ListenAndServe(":8080", nil)
}
