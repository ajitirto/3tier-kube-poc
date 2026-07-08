package main

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
)

const backendURL = "http://backend:8080"

func main() {

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		resp, err := http.Get(backendURL + "/api/visitors")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, string(body))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		resp, err := http.Get(backendURL + "/api/health")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()

		io.Copy(w, resp.Body)
	})

	http.HandleFunc("/visitors", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:

			resp, err := http.Get(backendURL + "/api/visitors")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			defer resp.Body.Close()

			io.Copy(w, resp.Body)

		case http.MethodPost:

			body, _ := io.ReadAll(r.Body)

			resp, err := http.Post(
				backendURL+"/api/visitors",
				"application/json",
				bytes.NewBuffer(body),
			)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			defer resp.Body.Close()

			io.Copy(w, resp.Body)

		case http.MethodDelete:

			req, _ := http.NewRequest(
				http.MethodDelete,
				backendURL+"/api/visitors",
				nil,
			)

			client := &http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			defer resp.Body.Close()

			io.Copy(w, resp.Body)

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
