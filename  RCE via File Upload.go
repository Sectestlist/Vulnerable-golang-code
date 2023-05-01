package main

import (
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseMultipartForm(10 << 20) // max memory limit
			file, _, err := r.FormFile("uploadedFile")
			if err != nil {
				http.Error(w, "Error retrieving file from form", http.StatusBadRequest)
				return
			}
			defer file.Close()

			// Executing the file directly without any validation
			cmd := exec.Command("./" + file.Name())
			err = cmd.Run()
			if err != nil {
				http.Error(w, "Error executing file", http.StatusInternalServerError)
				return
			}

			w.Write([]byte("File executed successfully!"))
		} else {
			// Render the upload form
			w.Write([]byte(`<form method="POST" action="/" enctype="multipart/form-data">Upload a file: <input type="file" name="uploadedFile"/><input type="submit"/></form>`))
		}
	})
	http.ListenAndServe(":8080", nil)
}
