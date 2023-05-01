package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file := r.URL.Query().Get("file")
		data, err := ioutil.ReadFile(file)
		if err != nil {
			http.Error(w, "Error: File not found", http.StatusNotFound)
			return
		}
		w.Write(data)
	})
	http.ListenAndServe(":8080", nil)
}
