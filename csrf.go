package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Perform some action without checking for CSRF token
			fmt.Fprint(w, "Action performed!")
		} else {
			// Render the form
			fmt.Fprint(w, `<form method="POST" action="/">Press me!<input type="submit"/></form>`)
		}
	})
	http.ListenAndServe(":8080", nil)
}
